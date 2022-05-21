package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
}

func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transaction []Transaction
	// Find all campaign based on `campaignID`
	err := r.db.Where("campaign_id = ?", campaignID).Preload("User").Order("id desc").Find(&transaction).Error
	// If error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
