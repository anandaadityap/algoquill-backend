package request

type RegisterUserInput struct {
	Name     string `json:"name" validate:"required,min=5,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
