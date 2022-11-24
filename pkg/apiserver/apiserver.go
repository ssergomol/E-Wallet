package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ssergomol/E-Wallet/pkg/database"
)

type APIserver struct {
	config *ConfigServer
	router *mux.Router
	logger *logrus.Logger
	db     *database.Storage
}

func CreateServer(config *ConfigServer) (*APIserver, error) {
	server := &APIserver{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	err := server.configureDatabase()
	if err != nil {
		return &APIserver{}, err
	}

	return server, nil
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	return nil
}

func (s *APIserver) configureRouter() {
	s.RegisterHome()
	s.RegisterBalance()
}

func (s *APIserver) configureDatabase() error {
	config, err := database.NewConfig()
	if err != nil {
		return err
	}

	db := database.NewDB(config)
	s.logger.Info("connecting to database")
	if err := db.Connect(); err != nil {
		return err
	}
	s.logger.Info("connection established")

	s.db = db
	return nil
}
