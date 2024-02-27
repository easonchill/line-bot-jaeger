package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MongoDB MongoDB
	Server  Server
	Line    Line
}

type MongoDB struct {
	UserName string
	Password string
	Host     string
	Database string
}

type Server struct {
	Port    string
	RunMode string
}

type Line struct {
	ChannelSecret string
	Token         string
}

func NewConfig() *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return &Config{MongoDB: MongoDB{UserName: viper.GetString("Database.UserName"),
		Password: viper.GetString("Database.PassWord"),
		Host:     viper.GetString("Database.Host"),
		Database: viper.GetString("Database.DBName")},
		Server: Server{Port: viper.GetString("Server.HttpPort")},
		Line:   Line{ChannelSecret: viper.GetString("Line.ChannelSecret"), Token: viper.GetString("Line.Token")},
	}
}
