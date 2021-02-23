package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"kougakusaischeduler/models"
)

func main() {
	DBMigrate(DBConnect())

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello())

	err := e.Start(getListeningPort())
	if err != nil {
		log.Fatalln(err)
	}
}

func getListeningPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8000"
}

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from echo.")
	}
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Event{})
	return db
}

func DBConnect() *gorm.DB {
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		dburl = "host=postgres port=5432 user=example_user dbname=example_db password=password sslmode=disable"
	}
	db, err := gorm.Open("postgres", dburl)
	if err != nil {
		panic(err)
	}
	return db
}
