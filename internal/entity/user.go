package entity

type Users struct {
	Username string `json:"username" bun:"username"`
	Password string `json:"password" bun:"password"`
	Role     string `json:"role" bun:"role"`
}

type CreateUserRequest struct {
	Username string `json:"username" bun:"id"`
	Password string `json:"password" bun:"password"`
	Role     string `json:"role" bun:"role"`
}

type CreateUserResponse struct {
	Id       int    `json:"id" bun:"id"`
	Username string `json:"username" bun:"username"`
	Password string `json:"password" bun:"password"`
	Role     string `json:"role" bun:"role"`
}

type UpdateUserRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUserColumnsRequest struct {
	ID     int               `json:"id"`
	Fields map[string]string `json:"fields"`
}

type UpdateUserResponse struct {
	Id       int     `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Role     *string `json:"role"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

type GetUserResponse struct {
	Id       int     `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Role     *string `json:"role"`
}

type Filter struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

type ListUserResponse struct {
	Users []*GetUserResponse `json:"users"`
	Total int                `json:"total"`
}
