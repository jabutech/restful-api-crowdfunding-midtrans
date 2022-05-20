package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
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
