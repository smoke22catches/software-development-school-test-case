package api

import (
	"github.com/gin-gonic/gin"

	"software-development-school-test-case/emails"
	"software-development-school-test-case/price"
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
	btcPrice, err := price.GetBtcPriceInUah()

	if err != nil {
		c.IndentedJSON(400, err.Error())
	}

	c.IndentedJSON(200, btcPrice)
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

func NotifySubscribersAboutBtcPrice(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	err := emails.SendBtcPriceToSubscribers()

	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, "Emails sent")
}

// structure of request to subscribe
type SubscribeEmailForm struct {
	Email string `form:"email" binding:"required,email"`
}
