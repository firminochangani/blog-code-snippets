package httpapi

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"apifirst/app"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
)

type Server struct {
	logger     *slog.Logger
	stdHttpSrv http.Server
}

func NewServer(ctx context.Context, port uint, application *app.App, logger *slog.Logger) (*Server, error) {
	router := echo.New()
	routerHandlers := &handlers{
		application: application,
	}

	apiSpec, err := GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("error loading the API specification: %v", err)
	}

	// Register middlewares to the router
	registerMiddlewares(router, apiSpec)

	// Register each handler to the router
	RegisterHandlers(router, routerHandlers)

	return &Server{
		logger: logger,
		stdHttpSrv: http.Server{
			Addr:              fmt.Sprintf(":%d", port),
			Handler:           router,
			ReadTimeout:       time.Second * 10,
			ReadHeaderTimeout: time.Second * 5,
			IdleTimeout:       time.Second * 60,
			BaseContext: func(listener net.Listener) context.Context {
				return ctx
			},
		},
	}, nil
}

func (s *Server) Start() error {
	s.logger.Info("The http server is about to run", "addr", s.stdHttpSrv.Addr)
	return s.stdHttpSrv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.stdHttpSrv.Shutdown(ctx)
}

func registerMiddlewares(router *echo.Echo, apiSpec *openapi3.T) {
	apiSpec.Servers = nil

	router.Use(middleware.Recover())
	router.Use(middleware.Logger())
	router.Use(oapimiddleware.OapiRequestValidatorWithOptions(apiSpec, &oapimiddleware.Options{}))
}
