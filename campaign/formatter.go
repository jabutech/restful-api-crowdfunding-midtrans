package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
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
	campaignFormatter.Slug = campaign.Slug
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
	/* - Create a var campaignsFormatter
	- with data type slice of struct CampaignFormatter (Note: Symbol {} this means is default value `empty array`)
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

// Struct CampaignDetailFormatter struct
type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	Slug             string                   `json:"slug"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"CurrentAmount"`
	BackerCount      int                      `json:"backer_count"`
	UserID           int                      `json:"user_id"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"` // Use type struct CampaignUserFormatter which in below
	Images           []CampaignImageFormatter `json:"images"`
}

// Struct for Campaign user formatter
type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// Struct for Campaign Image formatter
type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	// Create new object use CampaignDetailFormatter
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.BackerCount = campaign.BackerCount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.Slug = campaign.Slug

	// Handle image
	campaignDetailFormatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	// Handle property perks
	// Create var perks with value slice string
	var perks []string
	// Split perks and loop
	for _, perk := range strings.Split(campaign.Perks, ",") {
		// Insert every perks to var, and remove white space with package `strings.TrimSpace`
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	// Handle property user
	// Create new var object campaign user formatter
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = campaign.User.Name
	campaignUserFormatter.ImageUrl = campaign.User.AvatarFileName

	// Insert object campaignUserFormatter to struct CampaignDetailFormatter.User
	campaignDetailFormatter.User = campaignUserFormatter

	// Handle property images
	// Create var images with value slice of struct CampaignImageFormatter
	images := []CampaignImageFormatter{}
	// Loop campaign.CampaignImages
	for _, image := range campaign.CampaignImages {
		// Create new var object
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		// Create var isPrimary with value false
		isPrimary := false

		// Check if image.IsPrimary value is 1
		if image.IsPrimary == 1 {
			// Change value isPrimary to `true`
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary
		// Insert image to slice campaignImageFormatter
		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
