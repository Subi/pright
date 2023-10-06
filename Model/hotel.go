package model

import "time"

type Hotel struct {
	Id        int `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Location  string
	Price     float64
	Image     string
	Stars     float64
	Slug      string
}
