package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Contract service
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

// Struct service
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// Function service for register user
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

// Function service for handle login
func (s *service) Login(input LoginInput) (User, error) {
	// Get email and password from input user
	email := input.Email
	password := input.Password

	// Find user with repository FindByEmail
	user, err := s.repository.FindByEmail(email)
	// Check if error
	if err != nil {
		return user, err
	}

	// If user id is equal to 0 (not found)
	if user.ID == 0 {
		return user, errors.New("No user found on that email!")
	}

	// If user is available, compare password hash with password from request use bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	// Check if error
	if err != nil {
		return user, err
	}

	// If no error, return user
	return user, nil
}
