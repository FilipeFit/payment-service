package api

type GetAccountResponse struct {
	ID             uint64  `json:"id"`
	CustomerId     uint64  `json:"customerId"`
	CustomerName   string  `json:"customerName"`
	CustomerEmail  string  `json:"customer_email"`
	Balance        float64 `json:"balance"`
	Active         bool    `json:"active"`
	AllowOverdraft bool    `json:"allowOverdraft"`
}

type PaymentRequest struct {
	AccountID uint64  `json:"accountId"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
}

type PaymentResponse struct {
	AccountID uint64  `json:"accountId"`
	Balance   float64 `json:"balance"`
}

type ErrorResponse struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
}
