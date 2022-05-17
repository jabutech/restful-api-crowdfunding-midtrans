package user

// Struct for map user input register
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`        // validation has been required
	Occupation string `json:"occupation" binding:"required"`  // validation has been required
	Email      string `json:"email" binding:"required,email"` // validation has been required and format must been email
	Password   string `json:"password" binding:"required"`    // validation has been required
}

// Struct for map user input login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Struct for check email
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
