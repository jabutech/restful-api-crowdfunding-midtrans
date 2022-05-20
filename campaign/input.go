package campaign

// Struct for handle id campaign parameter from url
type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
