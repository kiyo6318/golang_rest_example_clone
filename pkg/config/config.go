package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Server *serverConfig
	DB		 *databaseConfig
}

type serverConfig struct {
	Port 						int 	 `env:"PORT,default=8080"`
	AllowCorsOrigin string `env:"ALLOW_CORS_ORIGIN,default=*"`
}