package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppTimeOut time.Duration

	DBMS       string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
}

func LoadConf() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	timeOutStr := os.Getenv("TIMEOUT_SECOND")
	timeOut, err := strconv.Atoi(timeOutStr)
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		AppTimeOut: time.Duration(timeOut) * time.Second,
		DBMS:       os.Getenv("DBMS"),
		DBPassword: os.Getenv("MYSQL_PASSWORD"),
		DBHost:     os.Getenv("MYSQL_DBHOST"),
		DBPort:     os.Getenv("MYSQL_DBPORT"),
		DBName:     os.Getenv("MYSQL_DATABASE"),
		DBUser:     os.Getenv("MYSQL_USER"),
	}
}
