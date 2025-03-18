package core

import (
	"github.com/gin-gonic/gin"
	"github.com/shplume/Modou/pkg/config"
)

type Server struct {
	Config config.ConfigReader
	Router *gin.Engine
}

var ServerInstance = NewDefaultServer()

func NewDefaultServer() *Server {
	return &Server{
		Config: config.NewDefaultConfigReader(),
	}
}

func (s *Server) Init() {
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

	s.Router = gin.Default()
}

func (s *Server) Run() {
	s.Init()

	s.Router.Run(s.Config.Get("address"))
}
