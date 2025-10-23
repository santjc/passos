package container

import (
	"net/http"
	"os"
	"strconv"

	"passos/internal/database"
	httplayer "passos/internal/http"
	pkgerrors "passos/internal/pkg/errors"
	"passos/internal/pkg/logger"
	"passos/internal/service"
)

type Container struct {
	logger logger.Logger
	db     database.Service
	port   int
	health service.HealthStatusProvider
}

func New() (*Container, error) {
	log := logger.New()
	db := database.New()

	port, err := loadPort()
	if err != nil {
		return nil, pkgerrors.Wrap("container.New", err, "unable to read PORT value")
	}

	return &Container{
		logger: log,
		db:     db,
		port:   port,
		health: service.NewHealthService(db),
	}, nil
}

func (c *Container) HTTPServer() *http.Server {
	cfg := httplayer.Config{
		Port:          c.port,
		Logger:        c.logger,
		HealthService: c.health,
	}
	return httplayer.NewServer(cfg)
}

func (c *Container) Logger() logger.Logger {
	return c.logger
}

func (c *Container) Close() error {
	if c == nil || c.db == nil {
		return nil
	}
	if err := c.db.Close(); err != nil {
		return pkgerrors.Wrap("container.Close", err, "closing database connection")
	}
	return nil
}

func loadPort() (int, error) {
	value := os.Getenv("PORT")
	if value == "" {
		return 0, nil
	}
	port, err := strconv.Atoi(value)
	if err != nil {
		return 0, pkgerrors.Wrap("container.loadPort", err, "invalid port value")
	}
	if port < 0 {
		return 0, pkgerrors.New("container.loadPort", "port value must be positive")
	}
	return port, nil
}
