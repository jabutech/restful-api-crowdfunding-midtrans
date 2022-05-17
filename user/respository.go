package user

import "gorm.io/gorm"

// Kontrak repository
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
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

// Function for find user by email
func (r *repository) FindByEmail(email string) (User, error) {
	// Create var user with type struct user
	var user User

	// Find from database
	err := r.db.Where("email = ?", email).Find(&user).Error
	// If error
	if err != nil {
		// return error
		return user, err
	}

	// If success, return user
	return user, nil
}
