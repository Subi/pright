package pinger

import (
	"log"
	"net/http"
	model "pright/Model"
	util "pright/Util"
	"strings"
)

type Pinger struct{}

var webhook string = "https://hooks.slack.com/services/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"

func New() *Pinger {
	return &Pinger{}
}

func (p *Pinger) NewHotel(hotel model.Hotel) {
	payload := util.CreateNewHotelPayload(hotel)
	sendAlert(payload)
}

func (p *Pinger) UpdatedHotel(foundHotel *model.Hotel, price float64) {
	payload := util.UpdateHotelPricePayload(foundHotel, price)
	sendAlert(payload)
}

func sendAlert(payload string) {
	c := http.Client{}
	r, err := http.NewRequest(http.MethodPost, webhook, strings.NewReader(payload))
	if err != nil {
		log.Printf("Error creating alert request %s", err)
	}
	r.Header.Add("Content-type", "application/json")
	resp, err := c.Do(r)
	if err != nil {
		log.Printf("Error executing client request %s", err)
	}
	if resp.StatusCode != 200 {
		log.Printf("Error sending payload data %s", err)
	}
}
