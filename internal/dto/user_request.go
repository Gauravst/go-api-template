package dto

type CreateUserRequest struct {
	Name     *string `json:"name,omitempty"`
	Email    string  `json:"email" validate:"required,email"`
	Username *string `json:"username,omitempty"`
	Password string  `json:"password" validate:"required,min=6"`
}
