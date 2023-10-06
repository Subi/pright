package util

import (
	"fmt"
	model "pright/Model"
	types "pright/Types"
)

func FormatHotelData(data types.Hotel) model.Hotel {
	if len(data.Node.Rates) < 1 || len(data.Node.Rates[0].RateAmounts) < 1 {
		return model.Hotel{}
	}
	newHotel := model.Hotel{
		Name:     data.Node.Property.BasicInformation.Name,
		Location: fmt.Sprintf(" %v %v", data.Node.Property.BasicInformation.Lat, data.Node.Property.BasicInformation.Long),
		Price:    float64(data.Node.Rates[0].RateAmounts[0].AmountPlusMandatoryFees.Locale.Value) / 100,
		Image:    getImage(data),
		Stars:    data.Node.Property.Reviews.Stars.Count,
		Slug:     getSlug(data.Node.Property.Nickname),
	}
	return newHotel
}

func getImage(data types.Hotel) string {
	if len(data.Node.Property.Media.PrimaryImage.Edges) < 1 {
		return ""
	}
	return fmt.Sprintf("https://cache.marriott.com/%s", data.Node.Property.Media.PrimaryImage.Edges[0].Node.ImageUrls.Horizational)
}

func getSlug(nickname string) string {
	return fmt.Sprintf("https://www.marriott.com/en-us/hotels/%s/overview", nickname)
}
