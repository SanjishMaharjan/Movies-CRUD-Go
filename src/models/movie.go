package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Rating   int    `json:"rating"`
	Genre    string `json:"genre"`
}
