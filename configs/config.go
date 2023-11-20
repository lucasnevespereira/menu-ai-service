package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbUrl             string
	DbUsersCollection string
	DbMenusCollection string
	DbName            string
}

func Load() Config {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	c := Config{}
	c.DbUrl = viper.GetString("MONGO_URL")
	c.DbUsersCollection = "users"
	c.DbMenusCollection = "menus"
	c.DbName = "menuai"

	return c
}
