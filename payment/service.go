package payment

import (
	"bwacroudfunding/campaign"
	"bwacroudfunding/transaction"
	"bwacroudfunding/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
	transactionRepository transaction.Repository
	campaignRepository    campaign.Repository
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
	ProcessPAyment(input transaction.TransactonNotificationInput) error
}

func NewService(transactionRepository transaction.Repository, campaignRepository campaign.Repository) *service {
	return &service{transactionRepository, campaignRepository}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = ""
	midclient.ClientKey = ""
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
func (s *service) ProcessPAyment(input transaction.TransactonNotificationInput) error {
	// Get transaction id and convert type to int
	transaction_id, _ := strconv.Atoi(input.OrderID)

	// Get transaction by id
	transaction, err := s.transactionRepository.GetByID(transaction_id)
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
	updatedTransaction, err := s.transactionRepository.Update(transaction)
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

		// Update campaihn
		_, err := s.campaignRepository.Update(campaign)
		// If error
		if err != nil {
			return err
		}
	}

	return nil
}
