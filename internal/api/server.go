package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fernandormoraes/transaction-demo/internal/api/handlers"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	echo       *echo.Echo
	db         *gorm.DB
	config     *Config
	httpClient *http.Client
}

func NewServer(db *gorm.DB, config *Config, httpClient *http.Client) *Server {
	return &Server{echo: echo.New(), db: db, config: config, httpClient: httpClient}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr: s.config.SERVER_HOST,
	}

	go func() {
		if err := s.echo.StartServer(server); err != nil {
			logrus.Fatal("Error starting Server: ", err)
		}
	}()

	s.echo.Validator = &handlers.CreateTransactionValidator{Validator: validator.New()}

	if err := s.Router(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	logrus.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
