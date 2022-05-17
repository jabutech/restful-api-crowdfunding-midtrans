package user

import "gorm.io/gorm"

// Kontrak repository
type Repository interface {
	Save(user User) (User, error)
}

// Struct repository
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Function for save data user
func (r *repository) Save(user User) (User, error) {
	// (1) Create new user on db with user data from request
	err := r.db.Create(&user).Error
	// (2) Check if error
	if err != nil {
		return user, err
	}

	// (3) If create user success, return user and error nil
	return user, nil

}
