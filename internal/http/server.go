package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"passos/internal/pkg/logger"
	"passos/internal/service"
)

type Config struct {
	Port          int
	Logger        logger.Logger
	HealthService service.HealthStatusProvider
}

type Server struct {
	port          int
	logger        logger.Logger
	healthService service.HealthStatusProvider
}

func NewServer(cfg Config) *http.Server {
	port := cfg.Port
	if port == 0 {
		if value := os.Getenv("PORT"); value != "" {
			if envPort, err := strconv.Atoi(value); err == nil && envPort > 0 {
				port = envPort
			}
		}
	}
	if port == 0 {
		port = 8080
	}

	srv := &Server{
		port:          port,
		logger:        cfg.Logger,
		healthService: cfg.HealthService,
	}

	if srv.logger == nil {
		srv.logger = logger.New()
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", srv.port),
		Handler:      srv.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
