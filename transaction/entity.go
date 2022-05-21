package transaction

import (
	"bwacroudfunding/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignId int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User // Get data user from table users
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
