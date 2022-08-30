package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	NetworkLink string

	RPCPort string
}

func Load() Config {
	c := Config{}

	c.PostgresHost = cast.ToString(look("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(look("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(look("POSTGRES_DATABASE", "iman"))
	c.PostgresUser = cast.ToString(look("POSTGRES_USER", "najmiddin"))
	c.PostgresPassword = cast.ToString(look("POSTGRES_PASSWORD", "1234"))

	c.NetworkLink = cast.ToString(look(`NETWORK_LINK`, `https://gorest.co.in/public/v1/posts`))
	c.RPCPort = cast.ToString(look("RPC_PORT", ":8000"))

	return c
}

func look(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
