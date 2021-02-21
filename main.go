package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	name string
}

func main() {
	DBMigrate(DBConnect())

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello())

	err := e.Start(":8000")
	if err != nil {
		log.Fatalln(err)
	}
}

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from echo.")
	}
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}

func DBConnect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=example_user dbname=example_db password=password sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
