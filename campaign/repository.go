package campaign

import "gorm.io/gorm"

// Contract repository campaign
type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

// Instance repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Function for find all data campaigns
func (r *repository) FindAll() ([]Campaign, error) {
	// Create var with value struct Campaign
	var campaigns []Campaign

	// Find campaign on database
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	// If error
	if err != nil {
		return campaigns, err
	}

	// If no error, return all data campains
	return campaigns, nil
}

// Function for find campaing by user id
func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	// Create var with value struct Campaign
	var campaigns []Campaign

	// Find campaign on db based on user id and preload or load relation with is_primary value 1 from campaign images
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	// If error
	if err != nil {
		return campaigns, err
	}

	// If no error, return all data campains
	return campaigns, nil
}

// Function for find campaign by id
func (r *repository) FindByID(ID int) (Campaign, error) {
	// Create var campaign value struct Campaign
	var campaign Campaign
	// Find campaign by id and load relation with table users and campaign_images
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	// If error
	if err != nil {
		return campaign, err
	}

	// If no error, return all data campains
	return campaign, nil
}
