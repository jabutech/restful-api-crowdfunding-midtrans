package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

// Instace service
func NewService(repository Repository) *service {
	return &service{repository}
}

// Function for find campaign
func (s *service) FindCampaigns(userId int) ([]Campaign, error) {
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
