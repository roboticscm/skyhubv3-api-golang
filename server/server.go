package server

import (
	"backend/system/db"
	"backend/system/error"
	"backend/system/route"
	"backend/system/slog"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port         int
	ServerEngine *echo.Echo
}

func (s *Server) Init() {
	s.ServerEngine = echo.New()
	s.ServerEngine.HTTPErrorHandler = error.CustomHTTPErrorHandler
	s.ServerEngine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}

func (s *Server) RegisterRoute(b *db.Broker) {
	if s.ServerEngine == nil {
		slog.Fatal("You must init Server first")
	}
	rootRoute := route.RootRoute{}
	rootRoute.RegisterRoute(s.ServerEngine, b)
}

func (s *Server) Start() {
	if s.ServerEngine == nil {
		slog.Fatal("You must init Server first")
	}
	s.ServerEngine.Logger.Fatal(s.ServerEngine.Start(fmt.Sprintf(":%v", s.Port)))
}
