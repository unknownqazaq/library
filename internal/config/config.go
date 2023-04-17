package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	defaultHTTPPort               = "8080"
	defaultHTTPReadTimeout        = 15 * time.Second
	defaultHTTPWriteTimeout       = 15 * time.Second
	defaultHTTPIdleTimeout        = 60 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
)

type (
	Config struct {
		HTTP     HTTPConfig
		POSTGRES DatabaseConfig
	}

	HTTPConfig struct {
		Port               string
		Host               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		IdleTimeout        time.Duration
		MaxHeaderMegabytes int
	}

	ClientConfig struct {
		Endpoint string
		Username string
		Password string
	}

	DatabaseConfig struct {
		URL string
	}
)

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init() (cfg Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	httpConfig := HTTPConfig{
		Port:               defaultHTTPPort,
		ReadTimeout:        defaultHTTPReadTimeout,
		WriteTimeout:       defaultHTTPWriteTimeout,
		IdleTimeout:        defaultHTTPIdleTimeout,
		MaxHeaderMegabytes: defaultHTTPMaxHeaderMegabytes,
	}
	cfg.HTTP = httpConfig

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return
	}

	err = envconfig.Process("HTTP", &cfg.HTTP)
	if err != nil {
		return
	}

	err = envconfig.Process("POSTGRES", &cfg.POSTGRES)
	if err != nil {
		return
	}

	return
}
