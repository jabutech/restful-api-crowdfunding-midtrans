package transaction

import "bwacroudfunding/user"

// Struct for handle id campaign transaction parameter from url
type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
