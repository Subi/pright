package handler

import (
	"errors"
	"log"
	model "pright/Model"
	pinger "pright/Pinger"

	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) SaveHotel(hotel model.Hotel) {
	err := h.DB.Save(&hotel).Error
	if errors.Is(err, gorm.ErrInvalidValue) {
		log.Printf("Error saving %s", hotel.Name)
	}
	p := pinger.New()
	p.NewHotel(hotel)
	log.Printf("Hotel %s successfully saved", hotel.Name)
}

func (h *Handler) UpdateHotel(hotel model.Hotel) {
	var foundHotel = &model.Hotel{}
	err := h.DB.First(&foundHotel, "name = ?", hotel.Name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Hotel %s not found", hotel.Name)
		h.SaveHotel(hotel)
		return
	}
	if foundHotel.Price != hotel.Price {
		err := h.DB.Model(&hotel).Where("name = ?", hotel.Name).Update("price", hotel.Price).Error
		if errors.Is(err, gorm.ErrInvalidData) {
			log.Println(err)
			return
		}
		p := pinger.New()
		p.UpdatedHotel(foundHotel, hotel.Price)
	}

}
