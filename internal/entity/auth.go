package entity

type RegisterRequest struct {
	Username string `json:"username" from:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterResponse struct {
	ID           int    `json:"id" form:"id"`
	Username     string `json:"username" form:"username"`
	Role         string `json:"role" form:"role"`
	Status       bool   `json:"status" form:"status"`
	AccessToken  string `json:"access_token" form:"access_token"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	ID           int    `json:"id" form:"id"`
	Username     string `json:"username" form:"username"`
	Role         string `json:"role" form:"role"`
	Status       bool   `json:"status" form:"status"`
	AccessToken  string `json:"access_token" form:"access_token"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

type NewAccessTokenResponse struct {
	ID           int    `json:"id" form:"id"`
	Username     string `json:"username" form:"username"`
	Role         string `json:"role" form:"role"`
	Status       bool   `json:"status" form:"status"`
	AccessToken  string `json:"access_token" form:"access_token"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}
