package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/mrKrabsmr/test-task-verba-group/configs"
	"github.com/mrKrabsmr/test-task-verba-group/internal/api"
	"github.com/mrKrabsmr/test-task-verba-group/internal/app"
)

type APIServer struct {
	logger      *slog.Logger
	router      *http.ServeMux
	config      configs.Config
	application *app.App
}

func NewAPIServer(config configs.Config) *APIServer {
	level := slog.LevelInfo
	if config.Debug {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	return &APIServer{
		logger:      logger,
		router:      http.NewServeMux(),
		config:      config,
		application: app.New(config, logger),
	}
}

func (s *APIServer) configureRoutes() {
	if s.config.Version == 1 {
		api.ConfigureV1Routes(s.router, s.application)
		return
	}

	panic("incorrect version")
}

func (s *APIServer) MustRun(init bool) {
	if init {
		s.application.InitDB()
	}

	s.configureRoutes()

	go func() {
		time.Sleep(time.Second * 1)
		s.logger.Info(fmt.Sprintf("APISERVER IS RUNNING AT %s", s.config.Address))
	}()

	if err := http.ListenAndServe(s.config.Address, s.router); err != nil {
		panic(err)
	}

}
