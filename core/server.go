package core

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/Modou/middleware"
	"github.com/shplume/Modou/pkg/config"
	"github.com/shplume/Modou/pkg/logger"
)

type Server struct {
	Config config.ConfigReader
	Logger logger.Logger
	*gin.Engine
}

type ServerOptionFunc func(*Server)

func WithConfigReader(config config.ConfigReader) ServerOptionFunc {
	return func(s *Server) {
		s.Config = config
	}
}

func WithLogger(logger logger.Logger) ServerOptionFunc {
	return func(s *Server) {
		s.Logger = logger
	}
}

func NewServer(opts ...ServerOptionFunc) *Server {
	server := &Server{}

	for _, opt := range opts {
		opt(server)
	}

	server.init()

	return server
}

func (s *Server) init() {
	s.Config.SetConfigPath([]string{"./config/", "."})

	if err := s.Config.LoadEnvConfig(".env"); err != nil {
		panic(err)
	}

	if err := s.Config.LoadDefaultConfig("config_base"); err != nil {
		panic(err)
	}

	mode := s.Config.Get("MODE")
	if mode == "" {
		mode = "development"
	}

	if mode == "development" {
		gin.SetMode(gin.DebugMode)

		if err := s.Config.LoadConfig("config_dev"); err != nil {
			panic(err)
		}
	} else {
		gin.SetMode(gin.ReleaseMode)

		if err := s.Config.LoadConfig("config_prod"); err != nil {
			panic(err)
		}
	}

	s.Engine = gin.Default()
}

func (s *Server) Run() {
	s.Engine.Run(s.Config.Get("address"))
}

var ServerInstance = NewDefaultServer()

func NewDefaultServer() *Server {
	opts := []ServerOptionFunc{
		WithConfigReader(config.DefaultConfigReaderInstance),
		WithLogger(logger.DefaultLoggerInstance),
	}

	server := NewServer(opts...)

	manager := NewContextManager()
	manager.Use(middleware.NewLoggerProvider(logger.DefaultLoggerInstance))

	server.Use(manager.Build())

	return server
}
