package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Number string
	Name   string
	Email  string
	Grade  int
}

type Event struct {
	gorm.Model
	Name  string
	Place string
}
