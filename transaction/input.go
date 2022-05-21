package transaction

import "bwacroudfunding/user"

// Struct for handle id campaign transaction parameter from url
type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

// Struct for handle create campaign transaction input
type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       user.User
}

// Struct for handle notification from midtrans
type TransactonNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
