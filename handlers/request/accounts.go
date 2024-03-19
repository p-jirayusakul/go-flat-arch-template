package request

type RegisterRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6"`
}
