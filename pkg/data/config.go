package data

import (
	"fmt"
)

type DatabaseConf struct {
	Host       string `mapstructure:"POSTGRES_HOST"`
	Port       string `mapstructure:"POSTGRES_PORT"`
	User       string `mapstructure:"POSTGRES_USER"`
	Password   string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	LifeTime   int    `mapstructure:"POSTGRES_LIFETIME"`
	MaxConn    int    `mapstructure:"POSTGRES_MAX_CONN"`
	MaxIdle    int    `mapstructure:"POSTGRES_MAX_IDLE"`
	connection string
}

func (r *DatabaseConf) GetConnection() string {
	if r.connection != "" {
		return r.connection
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		r.Host, r.Port, r.User, r.Password, r.DBName)
}

func (r *DatabaseConf) SetConnection(con string) {
	r.connection = con
}
