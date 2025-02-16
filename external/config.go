package external

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerHost     string
	DatabaseConfig DatabaseConfig
	HttpConfig     HttpConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type HttpConfig struct {
	ServiceURL string
	Timeout    time.Duration
}

var (
	runOnce sync.Once
	config  Config
)

func GetConfig() Config {
	runOnce.Do(func() {
		cfg, err := initConfig()
		if err != nil {
			fmt.Println(context.Background(), err, "could not load usecase configuration")
		}
		config = Config{
			ServerHost: cfg.GetString("server.host"),
			DatabaseConfig: DatabaseConfig{
				Host:     cfg.GetString("database.host"),
				Port:     cfg.GetString("database.port"),
				User:     cfg.GetString("database.user"),
				Password: cfg.GetString("database.password"),
				DbName:   cfg.GetString("database.dbname"),
			},
			HttpConfig: HttpConfig{
				ServiceURL: cfg.GetString("http.service_url"),
				Timeout:    cfg.GetDuration("http.timeout"),
			},
		}
	})

	return config
}

func initConfig() (viper.Viper, error) {
	cfg := viper.New()
	var err error
	initDefaults(cfg)
	// workaround because viper does not resolve envs when unmarshalling
	for _, key := range cfg.AllKeys() {
		val := cfg.Get(key)
		cfg.Set(key, val)
	}
	return *cfg, err
}

func initDefaults(config *viper.Viper) {
	config.SetDefault("server.host", "0.0.0.0:8000")
	config.SetDefault("database.host", "postgres")
	config.SetDefault("database.port", "5432")
	config.SetDefault("database.user", "root")
	config.SetDefault("database.password", "root")
	config.SetDefault("database.dbname", "root")
	config.SetDefault("http.service_url", "http://localhost:8080")
	config.SetDefault("http.timeout", 5*time.Second)
}
