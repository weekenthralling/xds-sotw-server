package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type StoreConfig struct {
	Type         string         `mapstructure:"type"`
	Postgres     PostgresConfig `mapstructure:"postgres"`
	MaxOpenConns int            `mapstructure:"max_open_conns"`
	MaxIdleConns int            `mapstructure:"max_idle_conns"`
}

type PostgresConfig struct {
	DSN string `mapstructure:"dsn"`
}

type Config struct {
	Debug bool `mapstructure:"debug"`
	Port  int  `mapstructure:"port"`

	// xDS-SOTW-Service node id
	NodeId string `mapstructure:"node_id"`

	// storage
	Store StoreConfig `mapstructure:"store"`
}

func LoadConfig() *Config {

	viper.AutomaticEnv()

	viper.SetEnvPrefix("SERVICE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// set env default
	viper.SetDefault("port", 18000)
	viper.SetDefault("store.max_open_conns", 10)
	viper.SetDefault("store.max_idle_conns", 5)

	// bind env
	viper.BindEnv("debug")
	viper.BindEnv("port")
	viper.BindEnv("node_id")
	viper.BindEnv("store.type")
	viper.BindEnv("store.postgres.dsn")
	viper.BindEnv("store.max_open_conns")
	viper.BindEnv("store.max_idle_conns")

	var config Config
	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	log.Infof("Config: %+v", config)
	return &config
}
