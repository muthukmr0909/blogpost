package db_connection

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	logging "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql" // MySQL driver import
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "blogpost/model"
)

var (
	Db       *gorm.DB
	err      error
	Blog     model.BlogArticle
	Comments model.ArticleComments
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		logging.Error("Error while loading .env file")
		os.Exit(1)
	}

	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// create a connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, port, dbName)
	// Replace with your database connection details
	Db, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		logging.Error("Database connection error:", err)
		return
	}
	logging.Info("Database connected succesfully")

	Db.AutoMigrate(&Blog)
	Db.AutoMigrate(&Comments)

}
