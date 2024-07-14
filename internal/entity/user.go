package entity

import "time"

type Users struct {
	ID        *int       `json:"id" bun:"id"`
	Username  string     `json:"username" bun:"username"`
	Password  string     `json:"password" bun:"password"`
	Role      string     `json:"role" bun:"role"`
	Status    bool       `json:"status" bun:"status"`
	Refresh   *string    `json:"refresh" bun:"refresh"`
	CreatedBy *int       `json:"created_by" bun:"created_by"`
	UpdatedBy *int       `json:"updated_by" bun:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at" bun:"updated_at"`
}

type CreateUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Status    bool   `json:"status"`
	CreatedBy int    `json:"-"`
}

type CreateUserResponse struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Role     string  `json:"role"`
	Status   bool    `json:"status"`
	Refresh  *string `json:"refresh"`
}

type UpdateUserRequest struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Role      string  `json:"role"`
	Status    bool    `json:"status"`
	Refresh   *string `json:"refresh"`
	UpdatedBy int     `json:"-"`
}

type UpdateUserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
	Refresh  string `json:"refresh"`
}

type UpdateUserColumnsRequest struct {
	ID     int               `json:"id"`
	Fields map[string]string `json:"fields"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

type GetUserResponse struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Role     string  `json:"role"`
	Status   bool    `json:"status"`
	Refresh  *string `json:"refresh"`
}

type Filter struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

type ListUserResponse struct {
	Users []*GetUserResponse `json:"users"`
	Total int                `json:"total"`
}
