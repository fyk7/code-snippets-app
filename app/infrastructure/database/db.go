package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/fyk7/code-snippets-app/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) *gorm.DB {
	val := url.Values{}
	val.Add("charset", "utf8mb4")
	val.Add("parseTime", "true")
	connection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return dbConn
}
