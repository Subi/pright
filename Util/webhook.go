package util

import (
	"fmt"
	model "pright/Model"
)

func CreateNewHotelPayload(hotel model.Hotel) string {
	return fmt.Sprintf(`{
		"blocks": [
			{
				"type": "section",
				"block_id": "sectionBlockOnlyMrkdwn",
				"text": {
					"type": "mrkdwn",
					"text": "*%s* has been added"
				}
			},	
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "*<%s|%s>*\n★★★★★\n$%v per night\nStars: %v "
				},
				"accessory": {
					"type": "image",
					"image_url": "%s",
					"alt_text": "%s thumbnail"
				}
			},
			{
				"type": "context",
				"elements": [
					{
						"type": "image",
						"image_url": "https://api.slack.com/img/blocks/bkb_template_images/tripAgentLocationMarker.png",
						"alt_text": "Location Pin Icon"
					},
					{
						"type": "plain_text",
						"emoji": true,
						"text": "Location: %s"
					}
				]
			},
			{
				"type": "divider"
			}
		]
	}`, hotel.Name, hotel.Slug, hotel.Name, fmt.Sprintf("%g", hotel.Price), hotel.Stars, hotel.Image, hotel.Name, hotel.Location)
}

func UpdateHotelPricePayload(hotel *model.Hotel, price float64) string {
	return fmt.Sprintf(`{
		"blocks": [
			{
				"type": "section",
				"block_id": "sectionBlockOnlyMrkdwn",
				"text": {
					"type": "mrkdwn",
					"text": "*%s* price has changed ~%v~ -> %v"
				}
			},
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "*<%s|%s>*\n★★★★★\n$%v per night\nStars: %v "
				},
				"accessory": {
					"type": "image",
					"image_url": "%s",
					"alt_text": "%s thumbnail"
				}
			},
			{
				"type": "context",
				"elements": [
					{
						"type": "image",
						"image_url": "https://api.slack.com/img/blocks/bkb_template_images/tripAgentLocationMarker.png",
						"alt_text": "Location Pin Icon"
					},
					{
						"type": "plain_text",
						"emoji": true,
						"text": "Location: %s"
					}
				]
			},
			{
				"type": "divider"
			}
		]
	}`, hotel.Name, fmt.Sprintf("%g", hotel.Price), fmt.Sprintf("%g", price), hotel.Slug, hotel.Name, fmt.Sprintf("%g", price), hotel.Stars, hotel.Image, hotel.Name, hotel.Location)
}
