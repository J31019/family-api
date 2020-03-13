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
	Token  string `json:"token,omitempty"`
}

// RegisterRequest represents register request model.
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// RegisterResponse represents register response model.
type RegisterResponse struct {
	Result bool   `json:"result"`
	Error  string `json:"error,omitempty"`
}

// SubmitRequest represents submit request model.
type SubmitRequest struct {
	Code string `json:"code"`
}

// SubmitResponse represents submit response model.
type SubmitResponse struct {
	Result bool   `json:"result"`
	Error  string `json:"error,omitempty"`
	Token  string `json:"token,omitempty"`
}
