package request

type RegisterUserRequest struct {
	Name string `json:"name" validator:"required"`
	Email string `json:"email" validator:"required"`
	Password string `json:"password" validator:"required, min=6"`
}

type LoginUserRequest struct {
	Email string `json:"email" validator:"required"`
	Password string `json:"password" validator:"required, min=6"`
}

type LogoutRequest struct {
	Token string `json:"token" validator:"required"`
}