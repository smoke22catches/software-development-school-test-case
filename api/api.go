package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"software-development-school-test-case/emails"

	"github.com/gin-gonic/gin"
)

func StartApi() {
	router := gin.Default()
	router.GET("/api/rate", GetBtcPriceInUah)
	router.POST("/api/subscribe", AddEmailToSubscriptionList)
	router.POST("/api/sendEmails", NotifySubscribersAboutBtcPrice)
	router.Run("0.0.0.0:5000")
}

func GetBtcPriceInUah(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	var responseData BtcPriceResponse

	// Get request to coinbase api
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?base=BTC&currency=UAH")
	if err != nil {
		c.IndentedJSON(400, -1)
		return
	}

	// decode api response
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&responseData)
	if err != nil {
		c.IndentedJSON(400, -1)
		return
	}

	// get price from response
	price, err := strconv.ParseFloat(responseData.Data.Amount, 64)
	if err != nil {
		c.IndentedJSON(400, -1)
		return
	}

	c.IndentedJSON(200, price)
}

func AddEmailToSubscriptionList(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var form SubscribeEmailForm
	err := c.Bind(&form)

	if err != nil {
		c.IndentedJSON(400, err.Error())
		return
	}

	emailAdded, err := emails.AddEmailToSubscriptionList(form.Email)

	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	if !emailAdded {
		c.IndentedJSON(409, "Email already exists")
		return
	}

	c.IndentedJSON(200, "Email added")
}

func NotifySubscribersAboutBtcPrice(c *gin.Context) {}

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

// structure of request to subscribe
type SubscribeEmailForm struct {
	Email string `form:"email" binding:"required,email"`
}
