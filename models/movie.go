package models

import "time"

type Movie struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date" time_format:"2006-01-02"`
	Rating      float64   `json:"rating"`
	PosterPath  string    `json:"poster_path"`
}
