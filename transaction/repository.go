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
	GetByUserID(userID int) ([]Transaction, error)
	GetByID(UserID int) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
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

func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	var transactions []Transaction

	/*
		- Find Transaction based on current user id is logged in
		- Load relation to table campaign and relation from campaign to campaign images, load only campaign images of is_primary is `1`
	*/
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	// If error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetByID(userID int) (Transaction, error) {
	// Create var transaction value struct Transaction
	var transaction Transaction
	// Find transaction by id
	err := r.db.Where("id = ?", userID).Find(&transaction).Error
	// If error
	if err != nil {
		return transaction, err
	}

	// If no error, return all data campains
	return transaction, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	// Create transaction on db
	err := r.db.Create(&transaction).Error
	// If error
	if err != nil {
		return transaction, err
	}
	// Return transaction
	return transaction, err
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	// Create transaction on db
	err := r.db.Save(&transaction).Error
	// If error
	if err != nil {
		return transaction, err
	}
	// Return transaction
	return transaction, err
}
