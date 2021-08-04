package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/iotexproject/iotex-core/blockchain/genesis"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	// Default is the default config
	Default = Config{
		Server: Server{},
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
		GrpcAPIPort int `yaml:"grpcApiPort"`
		HTTPAPIPort int `yaml:"httpApiPort"`
	}
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Debug    bool   `yaml:"debug"`
	}
	Genesis struct {
		VoteWeightCalConsts genesis.VoteWeightCalConsts `yaml:"voteWeightCalConsts"`
	}
	Config struct {
		Server   Server   `yaml:"server"`
		Database Database `yaml:"database"`
		Genesis  Genesis  `yaml:"genesis"`
	}
)

func New(path string) (cfg *Config, err error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, errors.Wrap(err, "failed to read config content")
	}
	cfg = &Default
	if err = yaml.Unmarshal(body, cfg); err != nil {
		return cfg, errors.Wrap(err, "failed to unmarshal config to struct")
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
