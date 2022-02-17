package web

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Id    string `json:"id"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
