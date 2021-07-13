package repositories

import (
	"github.com/filipeFit/payment-service/domain"
	"github.com/filipeFit/payment-service/domain/api"
	"gorm.io/gorm"
)

type paymentRepository struct {
	DB *gorm.DB
}

type paymentRepositoryInterface interface {
	Create(request *api.CreatePaymentRequest) (*domain.Payment, error)
	FindByPaymentID(paymentID uint64) (*domain.Payment, error)
	FindByAccountID(accountID uint64) ([]domain.Payment, error)
}

var (
	_ paymentRepositoryInterface
)

func init() {
	_ = &paymentRepository{}
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{DB: db}
}

func (s *paymentRepository) Create(request *api.CreatePaymentRequest) (*domain.Payment, error) {
	payment := api.ToPayment(request)
	result := s.DB.Create(&payment)
	if result.Error != nil {
		return nil, result.Error
	}
	return payment, nil
}

func (s *paymentRepository) FindByPaymentID(paymentID uint64) (*domain.Payment, error) {
	var payment domain.Payment
	result := s.DB.First(&payment, "id = ?", paymentID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &payment, nil
}

func (s *paymentRepository) FindByAccountID(accountID uint64) ([]domain.Payment, error) {
	var payments []domain.Payment
	result := s.DB.Find(&payments, "account_id = ?", accountID)
	if result.Error != nil {
		return nil, result.Error
	}
	return payments, nil
}
