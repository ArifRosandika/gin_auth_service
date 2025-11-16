package response

type RegisterUserResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}