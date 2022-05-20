package config

import (
	"fmt"
)

type DBConfig struct {
	User       string
	DbPassword string
	Driver     string
	DbName     string
	Host       string
	Port       string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		User:       "postgres",
		DbPassword: "local",
		DbName:     "voiceout",
		Host:       "0.0.0.0",
		Port:       "5432",
	}
}

func (d *DBConfig) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		d.Host,
		d.Port,
		d.User,
		d.DbName,
		d.DbPassword,
	)
}
