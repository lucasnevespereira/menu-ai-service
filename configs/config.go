package configs

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DbUrl             string
	DbUsersCollection string
	DbMenusCollection string
}

func Load() Config {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	c := Config{}
	c.DbUrl = os.Getenv("MONGO_URL")
	c.DbUsersCollection = "users"
	c.DbMenusCollection = "menus"

	return c
}
