package campaign

import "bwacroudfunding/user"

// Struct for handle id campaign parameter from url
type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

// Struct for handle create campaign input payload
type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

// Struct for handle upload campaign image
type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}
