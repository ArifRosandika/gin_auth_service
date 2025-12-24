package response

type RegisterUserResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}