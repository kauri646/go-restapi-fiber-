package entity

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FilmId    uint      `json:"film_id"`
	Film Film `gorm:"foreginKey:FilmReference"`
	Comment   string    `json:"comment"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}