package monitor

import (
	"encoding/json"
	"log"
	api "pright/API"
	handlers "pright/Handler"
	types "pright/Types"
	util "pright/Util"
	"strings"
	"time"
)

type Monitor struct {
	Api     *api.Api
	Handler *handlers.Handler
}

const url string = "https://www.marriott.com/mi/query/phoenixShopLowestAvailableRatesByGeoQuery"
const payload string = `{
"operationName":"phoenixShopLowestAvailableRatesByGeoQuery",
"query":"query phoenixShopLowestAvailableRatesByGeoQuery($search: LowestAvailableRatesGeolocationSearchInput, $offset: Int, $limit: Int, $sort: LowestAvailableRatesSearchSort, $filter: [PropertyDescriptionType], $modes: [RateMode]) {\n  searchLowestAvailableRatesByGeolocation(\n    search: $search\n    offset: $offset\n    limit: $limit\n    sort: $sort\n  ) {\n    pageInfo {\n      ...phoenixShopFragmentPageInfo\n      __typename\n    }\n    total\n    edges {\n      node {\n        distance\n        ...phoenixShopFragmentLARProperty\n        ...phoenixShopFragmentLARRates\n        isSavedProperty\n        __typename\n      }\n      __typename\n    }\n    facets {\n      ...phoenixShopFragmentLARFacets\n      __typename\n    }\n    status {\n      ... on ResponseStatus {\n        code\n        httpStatus\n        messages {\n          user {\n            message\n            type\n            id\n            field\n            details\n            __typename\n          }\n          ops {\n            id\n            type\n            field\n            message\n            details\n            __typename\n          }\n          dev {\n            type\n            id\n            field\n            message\n            details\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      ... on ResponseSuccess {\n        ...phoenixShopFragmentResponseSuccess\n        __typename\n      }\n      ... on NoHotelsFoundInThisDestinationError {\n        ...phoenixShopFragmentDestinationError\n        __typename\n      }\n      __typename\n    }\n    searchCenter {\n      ...phoenixShopFragmentSearchCenter\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment phoenixShopFragmentLARFacets on LowestAvailableRatesFacet {\n  type {\n    code\n    label\n    count\n    description\n    enumCode\n    __typename\n  }\n  buckets {\n    ... on LowestAvailableRatesTermFacetBucket {\n      code\n      label\n      description\n      count\n      __typename\n    }\n    ... on LowestAvailableRatesRangeFacetBucket {\n      index\n      start\n      end\n      count\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment phoenixShopFragmentLARRates on LowestAvailableRates {\n  rates {\n    rateAmounts(modes: $modes) {\n      amount {\n        origin {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        locale {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        __typename\n      }\n      rateUnit {\n        code\n        label\n        __typename\n      }\n      totalAmount {\n        origin {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        locale {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        __typename\n      }\n      points\n      taxes {\n        origin {\n          value\n          currency\n          valueDecimalPoint\n          __typename\n        }\n        locale {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        __typename\n      }\n      fees {\n        origin {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        locale {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        __typename\n      }\n      amountPlusMandatoryFees {\n        locale {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        origin {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        __typename\n      }\n      mandatoryFees {\n        origin {\n          valueDecimalPoint\n          value\n          currency\n          __typename\n        }\n        locale {\n          currency\n          value\n          valueDecimalPoint\n          __typename\n        }\n        __typename\n      }\n      rateMode {\n        code\n        label\n        __typename\n      }\n      __typename\n    }\n    membersOnly\n    lengthOfStay\n    rateCategory {\n      type {\n        code\n        __typename\n      }\n      value\n      __typename\n    }\n    status {\n      code\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment phoenixShopFragmentLARProperty on LowestAvailableRates {\n  property {\n    id\n    basicInformation {\n      resort\n      isAdultsOnly\n      brand {\n        id\n        name\n        type\n        photos {\n          content {\n            alternateText\n            index\n            name\n            url\n            __typename\n          }\n          type {\n            code\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      name\n      descriptions(filter: $filter) {\n        text\n        localizedText {\n          sourceText\n          translatedText\n          __typename\n        }\n        type {\n          code\n          __typename\n        }\n        __typename\n      }\n      currency\n      isRecentlyRenovated\n      isFullyRenovated\n      hasRenovatedRooms\n      ... on HotelBasicInformation {\n        newLobby\n        hasRenovatedRooms\n        isFullyRenovated\n        isRecentlyRenovated\n        newProperty\n        __typename\n      }\n      newProperty\n      openingDate\n      latitude\n      longitude\n      bookable\n      hasUniquePropertyLogo\n      __typename\n    }\n    reviews {\n      stars {\n        count\n        __typename\n      }\n      numberOfReviews {\n        count\n        __typename\n      }\n      __typename\n    }\n    media {\n      primaryImage {\n        edges {\n          node {\n            alternateDescription\n            title\n            imageUrls {\n              wideHorizontal\n              square\n              classicHorizontal\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    ... on Hotel {\n      seoNickname\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment phoenixShopFragmentPageInfo on PageInfo {\n  hasNextPage\n  hasPreviousPage\n  previousOffset\n  currentOffset\n  nextOffset\n  __typename\n}\n\nfragment phoenixShopFragmentResponseSuccess on ResponseSuccess {\n  code\n  httpStatus\n  messages {\n    user {\n      message\n      type\n      id\n      field\n      details\n      __typename\n    }\n    ops {\n      id\n      type\n      field\n      message\n      details\n      __typename\n    }\n    dev {\n      type\n      id\n      field\n      message\n      details\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment phoenixShopFragmentDestinationError on NoHotelsFoundInThisDestinationError {\n  code\n  httpStatus\n  messages {\n    user {\n      message\n      type\n      id\n      field\n      details\n      __typename\n    }\n    ops {\n      id\n      type\n      field\n      message\n      details\n      __typename\n    }\n    dev {\n      type\n      id\n      field\n      message\n      details\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment phoenixShopFragmentSearchCenter on PropertySearchCenter {\n  latitude\n  longitude\n  name\n  address\n  __typename\n}\n",
"variables":{"search":{"latitude":40.7830603,"longitude":-73.9712488,"distance":80467.2,"options":{"startDate":"2023-11-09","endDate":"2023-11-13","childAges":[],"rateRequestTypes":[{"type":"CLUSTER","value":"MMP"}],"numberInParty":1,"quantity":1,"includeMandatoryFees":true,"customerId":"","includeTaxesAndFees":false,"includeUnavailableProperties":true},"facets":{"terms":[{"type":"BRANDS","dimensions":[]},{"type":"AMENITIES","dimensions":[]},{"type":"PROPERTY_TYPES","dimensions":[]},{"type":"ACTIVITIES","dimensions":[]},{"type":"CITIES","dimensions":[]},{"type":"STATES","dimensions":[]},{"type":"COUNTRIES","dimensions":[]},{"type":"HOTEL_SERVICE_TYPES","dimensions":[]},{"type":"MEETINGS_EVENTS","dimensions":[]},{"type":"TRANSPORTATION_TYPES","dimensions":[]},{"type":"LEISURE_REGIONS","dimensions":[]}],"ranges":[{"type":"PRICE","dimensions":[],"endpoints":["0","100","200","overflow"]},{"type":"DISTANCE","dimensions":[],"endpoints":["0","5000","15000","80000"]}]}},"limit":75,"offset":0,"sort":{"fields":[{"field":"DISTANCE","direction":"ASC"}]},"filter":["HOTEL_MARKETING_CAPTION","RESORT_FEE_DESCRIPTION","DESTINATION_FEE_DESCRIPTION"],"modes":["LOWEST_AVERAGE_NIGHTLY_RATE","POINTS_PER_UNIT"]}}
`

func NewMonitor(api *api.Api, handler *handlers.Handler) *Monitor {
	return &Monitor{
		Api:     api,
		Handler: handler,
	}
}

func (m *Monitor) Init() {
	tick := time.NewTicker(10 * time.Second)

	for range tick.C {
		m.getHotelData()
	}
}

func (m *Monitor) getHotelData() {
	headers := map[string]string{
		"User-Agent":                   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		"Apollographql-Client-Name":    "phoenix_shop",
		"Apollographql-Client-Version": "v1",
		"Content-Type":                 "application/json",
		"Graphql-Operation-Signature":  "a604930e65e83fcfc4580c17c123ef063e9d274032eb849ffb2fa92164372d70",
	}
	Hotels := types.HotelsResponse{}
	resp, err := m.Api.Request("POST", url, headers, strings.NewReader(payload))
	defer resp.Body.Close()
	if err != nil {
		log.Printf("Error occured fetching hotel data : %s \n", err)
		return
	}
	if resp.StatusCode != 200 {
		log.Printf("Response status: %v", resp.StatusCode)
		return
	}
	decodeErr := json.NewDecoder(resp.Body).Decode(&Hotels)
	if decodeErr != nil {
		log.Printf("Error decoding Hotels JSON Response %v", decodeErr)
		return
	}

	for _, hotel := range Hotels.Data.SearchResults.Hotels {
		newHotel := util.FormatHotelData(hotel)
		m.Handler.UpdateHotel(newHotel)
	}
}
