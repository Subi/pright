package types

type HotelsResponse struct {
	Data Data `json:"data,omitempty"`
}

type Data struct {
	SearchResults SearchResults `json:"searchLowestAvailableRatesByGeolocation,omitempty"`
}

type SearchResults struct {
	Hotels []Hotel `json:"edges,omitempty"`
	Total  int     `json:"total,omitempty"`
}

type Hotel struct {
	Node Node `json:"node,omitempty"`
}

type Node struct {
	Distance float64  `json:"distance,omitempty"`
	Property Property `json:"property,omitempty"`
	Rates    []Rate   `json:"rates,omitempty"`
}

type Rate struct {
	StayLength  int `json:"lengthOfStay,omitempty"`
	RateAmounts []struct {
		Amount struct {
			Locale Locale `json:"locale,omitempty"`
		} `json:"amount,omitempty"`
		AmountPlusMandatoryFees struct {
			Locale Locale `json:"locale,omitempty"`
		} `json:"amountPlusMandatoryFees,omitempty"`
	} `json:"rateAmounts,omitempty"`
}

type Locale struct {
	Value        int `json:"value,omitempty"`
	DecimalPoint int `json:"valueDecimalPoint,omitempty"`
}

type Property struct {
	BasicInformation struct {
		Bookable bool `json:"bookable,omitempty"`
		Brand    struct {
			Name   string  `json:"name,omitempty"`
			Photos []Photo `json:"photos,omitempty"`
		} `json:"brand,omitempty"`
		Type string  `json:"type,omitempty"`
		Name string  `json:"name,omitempty"`
		Lat  float64 `json:"latitude,omitempty"`
		Long float64 `json:"longitude,omitempty"`
	} `json:"basicInformation,omitempty"`
	Media struct {
		PrimaryImage struct {
			Edges []struct {
				Node struct {
					ImageUrls struct {
						Horizational string `json:"classicHorizontal,omitempty"`
					} `json:"imageUrls,omitempty"`
				} `json:"node,omitempty"`
			} `json:"edges,omitempty"`
		} `json:"primaryImage,omitempty"`
	} `json:"media,omitempty"`
	Reviews struct {
		Stars struct {
			Count float64 `json:"count,omitempty"`
		} `json:"stars,omitempty"`
	} `json:"reviews,omitempty"`
	Nickname string `json:"seoNickName,omitempty"`
}

type Photo struct {
	Content []struct {
		URL string `json:"url,omitempty"`
	} `json:"content,omitempty"`
}
