package transaction

import (
	"bwacroudfunding/campaign"
	"bwacroudfunding/payment"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}
	// Get transaction use repository
	transaction, err := s.repository.GetByCampaignID(input.ID)
	// If error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	// Get transaction based on user id is logged in
	transactions, err := s.repository.GetByUserID(userID)
	// If error
	if err != nil {
		return transactions, err
	}
	// If no error, return transactions
	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	// Create new object transaction and passing data input
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	// Save transaction to db
	newTransaction, err := s.repository.Save(transaction)
	// If error
	if err != nil {
		return newTransaction, err
	}

	// Mapping data to struct payment.Transaction
	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}
	// Passing data to transaction midtrans
	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	// If error
	if err != nil {
		return newTransaction, err
	}

	// Mapping payment url to NewTransaction
	newTransaction.PaymentURL = paymentURL
	// Update transaction to db for inser payment url
	newTransaction, err = s.repository.Update(newTransaction)
	// If error
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil

}
