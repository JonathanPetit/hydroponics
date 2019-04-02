package database

import (
	"encoding/json"
	"os"
	"strconv"
	"fmt"
)

type Config struct {
	DB DBConfig `json:"db"`
}

type DBConfig struct {
	Name   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}


func GetConfigFromFile(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Unable to open configuration file at %v", path)
	}

	decoder := json.NewDecoder(file)
	conf := getDefaultConfig()

	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Printf("Error while parsing configuration file:\n%v", err)
	}

	return conf
}

func GetDBConfigFromEnv() DBConfig {
	conf := getDefaultDBConfig()

	updateValueFromEnv(&conf.Name, "DB_NAME")
	updateValueFromEnv(&conf.User, "DB_USER")
	updateValueFromEnv(&conf.Password, "DB_PASSWORD")
	updateValueFromEnv(&conf.Host, "DB_HOST")

	if port := os.Getenv("DB_PORT"); port != "" {
		if num, err := strconv.Atoi(port); err != nil {
			conf.Port = num
		}
	}

	return conf
}

func updateValueFromEnv(val *string, key string) {
	if v := os.Getenv(key); v != "" {
		*val = v
	}
}

func getDefaultDBConfig() DBConfig {
	return DBConfig{
		"hydro",
		"dev",
		"",
		"localhost",
		5432,
	}
}

func getDefaultConfig() Config {
	return Config{getDefaultDBConfig()}
}