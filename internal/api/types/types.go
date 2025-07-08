package types

type LoginRequest struct {
	Username string `json:"username" binding:"required, min=3,max=32"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required, min=3,max=32"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type LoginResponse struct {
	Token string
}
