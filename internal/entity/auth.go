package entity

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	Status       bool   `json:"status"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	Status       bool   `json:"status"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type NewAccessTokenResponse struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	Status       bool   `json:"status"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
