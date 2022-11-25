package config

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/imdario/mergo"
	"github.com/iotexproject/iotex-core/blockchain/genesis"
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
	}
	Genesis struct {
		VoteWeightCalConsts genesis.VoteWeightCalConsts `yaml:"voteWeightCalConsts"`
	}
	Config struct {
		Server   Server   `yaml:"server"`
		Database Database `yaml:"database"`
		RPC      string   `yaml:"rpc" env:"CHAIN_GRPC_ENDPOINT"`
		LogPath  string   `yaml:"logPath" env:"LOG_PATH"`
		Genesis  Genesis  `yaml:"genesis"`
	}
)

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
