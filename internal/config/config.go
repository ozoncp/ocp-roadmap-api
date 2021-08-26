package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

const CONFIG_NAME = "config.yml"

var once sync.Once
var instance *Config

type Metrics struct {
	Port   int    `yaml:"port"`
	Handle string `yaml:"handle"`
}

type REST struct {
	Port     int    `yaml:"port"`
	Endpoint string `yaml:"endpoint"`
}

type GRPC struct {
	Port int `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db-name"`
	SSLMode  string `yaml:"ssl-mode"`
}

type Roadmap struct {
	ButchSize uint `yaml:"create-butch-size"`
}

type Config struct {
	Roadmap  Roadmap  `yaml:"roadmap"`
	Database Database `yaml:"database"`
	GRPC     GRPC     `yaml:"grpc"`
	REST     REST     `yaml:"rest"`
	Metrics  Metrics  `yaml:"metrics"`
}

func createConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}

		file, err := os.Open(path)
		if err != nil {
			log.Err(err).Msg("error open config file")
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Err(err).Msg("error while closing file")
			}
		}()

		decoder := yaml.NewDecoder(file)
		if err = decoder.Decode(&instance); err != nil {
			log.Err(err).Msg("error while decode config file")
		}
	})

	return instance
}

func InitConfig(path string) *Config {
	return createConfig(path)
}
