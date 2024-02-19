package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	config := os.Getenv("CONFIG_PATH")
	if config == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	//check if file exists
	if _, err := os.Stat(config); os.IsNotExist(err) {
		log.Fatalf("Config file %s does not exist", config)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(config, &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &cfg
}
