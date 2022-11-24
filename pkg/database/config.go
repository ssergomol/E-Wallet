package database

import (
	toml "github.com/pelletier/go-toml"
)

type ConfigDB struct {
	dbServer string
	host     string
	dbName   string
	sslMode  string
	dbURL    string
}

func NewConfig() (*ConfigDB, error) {
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		return &ConfigDB{}, err
	}

	dbServer := "host=" + config.Get("store.host").(string) + " user=" + config.Get("store.user").(string) +
		" password=" + config.Get("store.password").(string) + " sslmode=" + config.Get("store.sslmode").(string)
	fullURL := "host=" + config.Get("store.host").(string) + " user=" + config.Get("store.user").(string) +
		" password=" + config.Get("store.password").(string) + " dbname=" + config.Get("store.dbname").(string) + " sslmode=" + config.Get("store.sslmode").(string)

	return &ConfigDB{
		dbServer: dbServer,
		host:     config.Get("store.host").(string),
		dbName:   config.Get("store.dbname").(string),
		sslMode:  config.Get("store.sslmode").(string),
		dbURL:    fullURL,
	}, nil
}
