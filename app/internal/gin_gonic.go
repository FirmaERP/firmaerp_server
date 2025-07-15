package internal

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type GinGonic struct {
	Port   string
	Engine *gin.Engine
}

func NewGin() (*GinGonic, error) {
	port := os.Getenv("GIN_PORT")
	if port == "" {
		return nil, errors.New("the environment variable GIN_PORT is empty")
	}

	engine := gin.Default()
	engine.Use(cors.Default())

	return &GinGonic{
		Port:   port,
		Engine: engine,
	}, nil
}

func (gin *GinGonic) Init() error {
	err := gin.Engine.Run(fmt.Sprintf(":%s", gin.Port))
	if err != nil {
		return errors.Wrap(err, "error starting the server")
	}

	return nil
}
