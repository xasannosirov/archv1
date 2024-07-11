package user

type Filter struct {
	Search *string
	Limit  *int
	Offset *int
	Page   *int
}

type AdminGetUserListResponse struct {
	ID       int     `json:"id"`
	Username *string `json:"username"`
	Role     *string `json:"role"`
}

type AdminCreateUserRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type AdminCreateUserResponse struct {
	ID int `json:"id"`
}

type ModeratorCreateUserRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type ModeratorCreateUserResponse struct {
	ID int `json:"id"`
}
