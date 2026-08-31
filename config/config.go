package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/imdario/mergo"
	"github.com/iotexproject/iotex-core/v2/blockchain/genesis"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v2"
)

var (
	// Default is the default config
	Default = Config{
		Server: Server{
			GrpcAPIPort: 8888,
			HTTPAPIPort: 8889,
		},
		Database: Database{
			Driver: "postgres",
			Host:   "127.0.0.1",
			Port:   "5432",
			User:   "postgres",
			Name:   "test",
		},
		RPC: "api.iotex.one:443",
		Genesis: Genesis{
			VoteWeightCalConsts: genesis.VoteWeightCalConsts{
				DurationLg: 1.2,
				AutoStake:  1,
				SelfStake:  1.06,
			},
			// Copied from iotex-core's genesis rather than imported: the field
			// landed in v2.5.0 and this module is on v2.2.0, whose dependency
			// graph no longer resolves (erigon-lib moved) without the replace
			// directives iotex-analyser carries. Two string constants that
			// change roughly never are not worth pulling erigon, pebble and
			// sentry into a read-only API to reach.
			//
			// Source: blockchain/genesis/genesis.go, Rewarding.HermesRewardVaultAddresses.
			// Verified against chain data 2026-08-31: on mainnet these are the
			// reward addresses of 44 and 15 candidates respectively, and every
			// other candidate has one of its own.
			HermesRewardVaultAddresses: []string{
				"io19604a05s2p3mecam2zz7d27hcr6ndyw80wvkmh",
				"io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85",
			},
		},
	}
)

type (
	Server struct {
		GrpcAPIPort int `yaml:"grpcApiPort" env:"GRPC_API_PORT"`
		HTTPAPIPort int `yaml:"httpApiPort" env:"HTTP_API_PORT"`
	}
	Database struct {
		Driver   string `yaml:"driver" env:"DB_DRIVER"`
		Host     string `yaml:"host" env:"DB_HOST"`
		Port     string `yaml:"port" env:"DB_PORT"`
		User     string `yaml:"user"  env:"DB_USER"`
		Password string `yaml:"password"  env:"DB_PASSWORD"`
		Name     string `yaml:"name"  env:"DB_NAME"`
		Debug    bool   `yaml:"debug"  env:"DB_DEBUG"`
		// SkipAutoMigrate disables the GORM AutoMigrate run at startup. Required
		// when pointing the API at a read-only PG standby — AutoMigrate issues
		// DDL and a standby rejects writes with SQLSTATE 25006.
		SkipAutoMigrate bool `yaml:"skipAutoMigrate"  env:"DB_SKIP_AUTO_MIGRATE"`
	}
	Genesis struct {
		VoteWeightCalConsts genesis.VoteWeightCalConsts `yaml:"voteWeightCalConsts"`
		// HermesRewardVaultAddresses are the reward addresses whose delegates
		// the Hermes service pays out for. It is what separates a Hermes
		// delegate from one distributing rewards on its own terms -- both have
		// the IIP-59 opt-in bit clear, and only the reward address tells them
		// apart.
		//
		// Defaults to iotex-core's genesis list. Overridable because this API
		// carries no per-network genesis: config.Genesis holds a handful of
		// fields rather than the real thing, and common/epoch.go reads
		// genesis.Default outright. A network whose genesis sets a different
		// list therefore needs to say so here rather than have this answer
		// silently wrong.
		HermesRewardVaultAddresses []string `yaml:"hermesRewardVaultAddresses"`
	}
	Config struct {
		Server   Server   `yaml:"server"`
		Database Database `yaml:"database"`
		RPC      string   `yaml:"rpc" env:"CHAIN_GRPC_ENDPOINT"`
		// RPCInsecure dials the chain endpoint without TLS. Defaults to false
		// so production endpoints keep their transport credentials; a local
		// node serves plaintext gRPC and otherwise fails every chain-meta
		// query with "first record does not look like a TLS handshake".
		RPCInsecure        bool    `yaml:"rpcInsecure" env:"CHAIN_GRPC_INSECURE"`
		EthArchiveEndPoint string  `yaml:"ethArchiveEndPoint" env:"ETH_ARCHIVE_ENDPOINT"`
		LogPath            string  `yaml:"logPath" env:"LOG_PATH"`
		Genesis            Genesis `yaml:"genesis"`
	}
)

// String masks the DB password so accidentally logging the config (e.g.
// main.go's startup `loaded config: %+v` line) doesn't leak the plaintext
// credential into container stdout. The alias type prevents infinite
// recursion — fmt won't call String on the alias since it has no method set.
func (d Database) String() string {
	type alias Database
	cp := alias(d)
	if cp.Password != "" {
		cp.Password = "***"
	}
	return fmt.Sprintf("%+v", cp)
}

func New(path string) (cfg *Config, err error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, errors.Wrap(err, "failed to read config content")
	}
	cfg = &Default
	var envCfg Config
	if err := envconfig.Process(context.Background(), &envCfg); err != nil {
		return cfg, errors.Wrap(err, "failed to process envconfig to struct")
	}
	if err = yaml.Unmarshal(body, cfg); err != nil {
		return cfg, errors.Wrap(err, "failed to unmarshal config to struct")
	}
	if err := mergo.Merge(&Default, envCfg, mergo.WithOverride); err != nil {
		return cfg, errors.Wrap(err, "failed to merge config")
	}
	return
}

var (
	// File names from which we attempt to read configuration.
	DefaultConfigFiles = []string{"config.yml", "config.yaml"}

	// Launchd doesn't set root env variables, so there is default
	DefaultConfigDirs = []string{getCurrentDirectory(), "~/.iotex-analyser-api", "/usr/local/etc/iotex-analyser-api", "/etc/iotex-analyser-api"}
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// FileExists checks to see if a file exist at the provided path.
func FileExists(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// ignore missing files
			return false, nil
		}
		return false, err
	}
	f.Close()
	return true, nil
}

// FindDefaultConfigPath returns the first path that contains a config file.
// If none of the combination of DefaultConfigDirs and DefaultConfigFiles
// contains a config file, return empty string.
func FindDefaultConfigPath() string {
	for _, configDir := range DefaultConfigDirs {
		for _, configFile := range DefaultConfigFiles {
			dirPath, err := homedir.Expand(configDir)
			if err != nil {
				continue
			}
			path := filepath.Join(dirPath, configFile)
			if ok, _ := FileExists(path); ok {
				return path
			}
		}
	}
	return ""
}
