package api

import "github.com/filipeFit/payment-service/domain"

type CreatePaymentRequest struct {
	AccountId   uint64  `json:"accountId"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Channel     string  `json:"channel"`
	Description string  `json:"description"`
	Destination string  `json:"destination"`
}

type CreatePaymentResponse struct {
	ID             uint64  `json:"id"`
	AccountId      uint64  `json:"accountId"`
	Type           string  `json:"type"`
	Amount         float64 `json:"amount"`
	Channel        string  `json:"channel"`
	Description    string  `json:"description"`
	Destination    string  `json:"destination"`
	AccountBalance float64 `json:"accountBalance,omitempty"`
}

func ToPayment(request *CreatePaymentRequest) *domain.Payment {
	payment := domain.Payment{
		AccountId:   request.AccountId,
		Description: request.Description,
		Amount:      request.Amount,
		Channel:     request.Channel,
		Destination: request.Destination,
		Type:        request.Type,
	}
	return &payment
}

func ToPaymentResponse(payment *domain.Payment) *CreatePaymentResponse {
	paymentResponse := CreatePaymentResponse{
		AccountId:   payment.AccountId,
		Description: payment.Description,
		Amount:      payment.Amount,
		Channel:     payment.Channel,
		ID:          payment.ID,
		Type:        payment.Type,
		Destination: payment.Destination,
	}
	return &paymentResponse
}
