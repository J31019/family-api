package model

// LoginRequest represents login request model.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents login response model.
type LoginResponse struct {
	Result bool   `json:"result"`
	Error  string `json:"error,omitempty"`
}
