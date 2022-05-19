package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

// Function for convert struct for handle single data campaign to CampaignFormatter
func FormatCampaign(campaign Campaign) CampaignFormatter {
	// Create new object
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount

	// Handle image
	campaignFormatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

// Function for convert struct for handle multiple data campaign to CampaignFormatter
func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	/* - Create a var campaignsFormatter with data type slice of struct CampaignFormatter (Symbol {} this means is default value `empty array`)
	- for accomodate multiple data campaign
	*/
	campaignsFormatter := []CampaignFormatter{}

	// Do loop multiple campaigns
	for _, campaign := range campaigns {
		// Create new object from every data campaign with function for handle Single FormatCampaign
		campaignFormatter := FormatCampaign(campaign)
		// Insert every data campaign to var campaignsFormatter
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	// Return var campaignsFormatter which there is data many campaigns
	return campaignsFormatter
}
