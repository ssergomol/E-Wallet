package apiserver

import (
	"github.com/pelletier/go-toml"
)

type ConfigServer struct {
	BindAddress string
	DatabaseURL string
	LogLevel    string
}

func NewConfig() (*ConfigServer, error) {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		return &ConfigServer{}, err
	}

	params := "host=" + config.Get("store.host").(string) + " dbname=" +
		config.Get("store.dbname").(string) + " sslmode=" + config.Get("store.sslmode").(string)
	return &ConfigServer{
		BindAddress: config.Get("bind_addr").(string),
		LogLevel:    config.Get("log_level").(string),
		DatabaseURL: params,
	}, nil
}
