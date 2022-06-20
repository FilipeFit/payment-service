package services

import (
	"github.com/filipeFit/payment-service/client"
	"github.com/filipeFit/payment-service/config"
	"github.com/filipeFit/payment-service/domain/api"
	"github.com/filipeFit/payment-service/handlers"
	"github.com/filipeFit/payment-service/repositories"
	"net/http"
)

type paymentService struct{}

type paymentServiceInterface interface {
	CreatePayment(request *api.CreatePaymentRequest, authorization string) (*api.CreatePaymentResponse, handlers.ApiError)
	FindByAccountID(accountID uint64) ([]api.CreatePaymentResponse, handlers.ApiError)
}

var (
	PaymentService    paymentServiceInterface
	paymentRepository = repositories.NewPaymentRepository(config.DB)
)

func init() {
	PaymentService = &paymentService{}
}

func (s *paymentService) CreatePayment(request *api.CreatePaymentRequest, authorization string) (*api.CreatePaymentResponse, handlers.ApiError) {
	accountCreatePayment := api.PaymentRequest{
		AccountID: request.AccountId,
		Type:      request.Type,
		Amount:    request.Amount,
	}
	accountBalance, err := client.ChangeAccountBalance(&accountCreatePayment, authorization)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusBadRequest, err.Error())
	}

	payment, err := paymentRepository.Create(request)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError, "error in creating the payment")
	}
	paymentResponse := api.CreatePaymentResponse{
		AccountId:      request.AccountId,
		Type:           request.Type,
		Amount:         request.Amount,
		Destination:    request.Destination,
		Channel:        request.Channel,
		ID:             payment.ID,
		Description:    request.Description,
		AccountBalance: accountBalance.Balance,
	}
	return &paymentResponse, nil
}

func (s *paymentService) FindByAccountID(accountID uint64) ([]api.CreatePaymentResponse, handlers.ApiError) {
	payments, err := paymentRepository.FindByAccountID(accountID)
	if err != nil {
		return nil, handlers.NewApiError(http.StatusInternalServerError, "error searching for payments")
	}

	var paymentsResponse []api.CreatePaymentResponse
	for _, payment := range payments {
		paymentResponse := api.ToPaymentResponse(&payment)
		paymentsResponse = append(paymentsResponse, *paymentResponse)
	}
	return paymentsResponse, nil
}
