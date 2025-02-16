package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Completed bool   `json:"completed"`
	Title     string `json: "title"`
}
