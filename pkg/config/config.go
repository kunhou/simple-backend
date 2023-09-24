package config

import (
	"fmt"

	"github.com/spf13/viper"

	"github/kunhou/simple-backend/pkg/data"
	http "github/kunhou/simple-backend/pkg/servmanager/http"
)

type AllConfig struct {
	Debug  bool
	Data   Data   `mapstructure:",squash"`
	Server Server `mapstructure:",squash"`
}

type Data struct {
	Database data.DatabaseConf `mapstructure:",squash"`
}

type Server struct {
	HTTP http.Config `mapstructure:",squash"`
}

func ReadConfig() (c *AllConfig, err error) {
	c = &AllConfig{}
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// parse config
	if err := viper.Unmarshal(c); err != nil {
		return c, err
	}

	return
}
