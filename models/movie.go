package models

import "gorm.io/gorm"

type Movie struct {
    gorm.Model
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}