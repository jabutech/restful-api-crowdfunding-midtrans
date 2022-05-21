package transaction

import (
	"bwacroudfunding/campaign"
	"bwacroudfunding/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User         // Get data user from table users
	Campaign   campaign.Campaign // Get data campaign from table campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
