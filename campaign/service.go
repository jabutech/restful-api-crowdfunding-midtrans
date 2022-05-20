package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(campaignID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

// Instace service
func NewService(repository Repository) *service {
	return &service{repository}
}

// Function for find campaign
func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	// If value userID is not 0
	if userId != 0 {
		// Find campaign by user id
		campaigns, err := s.repository.FindByUserID(userId)
		// If error
		if err != nil {
			return campaigns, err
		}

		// If no error, return campaigns
		return campaigns, nil
	}

	// Find all campaing
	campaigns, err := s.repository.FindAll()
	// If error
	if err != nil {
		return campaigns, err
	}

	// If no error, return campaigns
	return campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	// Find campaign by id use repository
	campaign, err := s.repository.FindByID(input.ID)
	// If error
	if err != nil {
		return campaign, err
	}

	// If no error, return campaign
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	// Create new object Campaign
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	// Create format slug name (campaign name + id user)
	slugName := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	// Make slug and pass to the object campaign
	campaign.Slug = slug.Make(slugName)

	// Save campaign to db with repository
	newCampaign, err := s.repository.Save(campaign)
	// Check if error
	if err != nil {
		return newCampaign, err
	}

	// If no error, return new campaign
	return newCampaign, nil
}

func (s *service) UpdateCampaign(campaignID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	// Find campaign by id
	campaign, err := s.repository.FindByID(campaignID.ID)
	// If error
	if err != nil {
		return campaign, err
	}

	// Check if campaign user not same with current user login which update campaign
	if campaign.UserID != inputData.User.ID {
		// Return response error
		return campaign, errors.New("Not a owner of the campaign.")
	}

	// Passing data payload input in to object campaign
	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	// Update data campaign
	updatedCampaign, err := s.repository.Update(campaign)
	// If error
	if err != nil {
		return updatedCampaign, err
	}

	// If success, return updatedCampaign
	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	// Find campaign by id
	campaign, err := s.repository.FindByID(input.CampaignID)
	// If error
	if err != nil {
		return CampaignImage{}, err
	}

	// Check if campaign user not same with current user login which update campaign
	if campaign.UserID != input.User.ID {
		// Return response error
		return CampaignImage{}, errors.New("Not a owner of the campaign.")
	}

	isPrimary := 0
	// If value `input.IsPrimary` is true
	if input.IsPrimary {
		// Change value var `isPrimary` is 1 (true)
		isPrimary = 1
		// Change value all columns `is_primary` to `0`
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		// If error
		if err != nil {
			return CampaignImage{}, err
		}
	}

	// Create new object
	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	// Save image to db
	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	// If error
	if err != nil {
		return newCampaignImage, err
	}

	// If success, return data newCampaignImage
	return newCampaignImage, nil

}
