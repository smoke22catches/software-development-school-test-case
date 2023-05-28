package price

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBtcPriceInUah() (float64, error) {
	var responseData BtcPriceResponse

	// Get request to coinbase api
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?base=BTC&currency=UAH")
	if err != nil {
		return -1, err
	}

	// decode api response
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&responseData)
	if err != nil {
		return -1, err
	}

	// get price from response
	price, err := strconv.ParseFloat(responseData.Data.Amount, 64)
	if err != nil {
		return -1, err
	}

	return price, nil
}

// structure of response from coinbase api
type BtcPriceResponse struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
}

func (r *BtcPriceResponse) UnmarshalJSON(data []byte) error {
	var response interface{}
	err := json.Unmarshal(data, &response)

	if err != nil {
		return err
	}

	responseMap := response.(map[string]interface{})
	responseData := responseMap["data"].(map[string]interface{})

	r.Data.Base = responseData["base"].(string)
	r.Data.Currency = responseData["currency"].(string)
	r.Data.Amount = responseData["amount"].(string)

	return nil
}
