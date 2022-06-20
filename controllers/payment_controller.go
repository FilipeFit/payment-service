package controllers

import (
	"github.com/filipeFit/payment-service/domain/api"
	"github.com/filipeFit/payment-service/handlers"
	"github.com/filipeFit/payment-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	paymentService = services.PaymentService
)

func CreatePayment(c *gin.Context) {
	var request api.CreatePaymentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := handlers.NewBadRequestError("invalid json body")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	authorization := c.Request.Header.Get("Authorization")
	response, err := paymentService.CreatePayment(&request, authorization)
	if err != nil {
		c.JSON(err.ResponseStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func FindPaymentsByAccountID(c *gin.Context) {
	accountId, err := strconv.ParseUint(c.Param("accountId"), 10, 64)
	if err != nil {
		apiErr := handlers.NewBadRequestError("invalid account id")
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	response, apiErr := paymentService.FindByAccountID(accountId)
	if apiErr != nil {
		c.JSON(apiErr.ResponseStatus(), apiErr)
		return
	}
	c.JSON(http.StatusOK, response)
}
