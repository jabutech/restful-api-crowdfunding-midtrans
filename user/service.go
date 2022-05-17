package user

import "golang.org/x/crypto/bcrypt"

// Contract service
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

// Struct service
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// Function for register user
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	// Create new object User from parameter input
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.Role = "user"

	// Hash password input with package bcrypy
	passworHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	// Check if bcrypt error
	if err != nil {
		return user, err
	}
	// Insert password hash to struct PasswordHash
	user.PasswordHash = string(passworHash)

	// Save user to db with repository
	newUser, err := s.repository.Save(user)
	// Check if error
	if err != nil {
		return newUser, err
	}

	// If success, return new user
	return newUser, nil

}
