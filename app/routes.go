package app

import (
	"github.com/filipeFit/payment-service/controllers"
	"github.com/filipeFit/payment-service/controllers/health"
)

func mapRoutes() {
	router.GET("/health", health.HealthCheck)
	router.POST("/payments", controllers.CreatePayment)
	router.GET("/payments/:accountId", controllers.FindPaymentsByAccountID)
}
