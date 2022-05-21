package transaction

import (
	"bwacroudfunding/campaign"
	"bwacroudfunding/payment"
	"errors"
	"strconv"
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
	ProcessPayment(input TransactonNotificationInput) error
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

func (s *service) ProcessPayment(input TransactonNotificationInput) error {
	// Get transaction id and convert type to int
	transaction_id, _ := strconv.Atoi(input.OrderID)

	// Get transaction by id
	transaction, err := s.repository.GetByID(transaction_id)
	// If error
	if err != nil {
		return err
	}

	// If no error, do if check status from midtrans
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	// Update transaction with appropriate status
	updatedTransaction, err := s.repository.Update(transaction)
	// If error
	if err != nil {
		return err
	}

	// Find campaign by id
	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
	// If error
	if err != nil {
		return err
	}

	// If status transaction is paid
	if updatedTransaction.Status == "paid" {
		// Update backer count added 1
		campaign.BackerCount = campaign.BackerCount + 1
		// Update current amount with updatedTransaction.Amount
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		// Update campaign
		_, err := s.campaignRepository.Update(campaign)
		// If error
		if err != nil {
			return err
		}
	}

	return nil
}
