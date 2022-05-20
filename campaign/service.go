package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
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
