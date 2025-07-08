package types

type LoginRequest struct {
	Password string `json:"password" binding:"required,min=8,max=64"`
	Username string `json:"username" binding:"required, min=3,max=32"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
	Username string `json:"username" binding:"required, min=3,max=32"`
}

type LoginResponse struct {
	Token string
}
