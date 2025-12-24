package request

type AuthRequest struct {
	RefreshToken string `json:"refresh_token"`
}