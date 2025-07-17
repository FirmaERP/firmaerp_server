package main

import (
	"os"
	"time"

	"github.com/FirmaERP/firmaerp_server/internal"
)

func main() {
	time.Local = time.UTC

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		panic("environment variable GIN_MODE is empty")
	}

	_, err := internal.InitGorm()
	if err != nil {
		panic(err)
	}

	gin, err := internal.NewGin()
	if err != nil {
		panic(err)
	}

	if mode != "release" {
		internal.InitOpenAPI(gin)
	}

	err = gin.Init()
	if err != nil {
		panic(err)
	}
}
