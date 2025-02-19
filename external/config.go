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
				Host:     cfg.GetString("DATABASE_HOST"),
				Port:     cfg.GetString("DATABASE_PORT"),
				User:     cfg.GetString("DATABASE_USER"),
				Password: cfg.GetString("DATABASE_PASSWORD"),
				DbName:   cfg.GetString("DATABASE_DBNAME"),
			},
			HttpConfig: HttpConfig{
				ServiceURL: cfg.GetString("FASTFOOD_PAYMENT_APP_URL"),
				Timeout:    cfg.GetDuration("HTTP_TIMEOUT"),
			},
		}
	})

	return config
}

func initConfig() (viper.Viper, error) {
	cfg := viper.New()
	var err error

	// workaround because viper does not resolve envs when unmarshalling
	for _, key := range cfg.AllKeys() {
		val := cfg.Get(key)
		cfg.Set(key, val)
	}
	initDefaults(cfg)
	fmt.Println(cfg)
	return *cfg, err
}

func initDefaults(config *viper.Viper) {
	config.SetDefault("server.host", "0.0.0.0:8000")
	config.SetDefault("DATABASE_HOST", "orders-db.cemq7svvd1t8.us-east-1.rds.amazonaws.com")
	config.SetDefault("DATABASE_PORT", "5432")
	config.SetDefault("DATABASE_USER", "postgres")
	config.SetDefault("DATABASE_PASSWORD", "postgres")
	config.SetDefault("DATABASE_DBNAME", "orders")
	config.SetDefault("FASTFOOD_PAYMENT_APP_URL", "http://a71651bd431364e399616a6c8cb93a80-882634322.us-east-1.elb.amazonaws.com:8000")
	config.SetDefault("HTTP_TIMEOUT", 5*time.Second)
}
