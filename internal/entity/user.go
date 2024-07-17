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
	Username  string `json:"username" xml:"username" yaml:"username" toml:"username" form:"username" query:"username"`
	Password  string `json:"password" xml:"password" yaml:"password" toml:"password" form:"password" query:"password"`
	Role      string `json:"role" xml:"role" yaml:"role" toml:"role" form:"role" query:"role"`
	Status    bool   `json:"status" xml:"status" yaml:"status" toml:"status" form:"status" query:"status"`
	CreatedBy int    `json:"-" bun:"created_by"`
}

type CreateUserResponse struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Role     string  `json:"role"`
	Status   bool    `json:"status"`
	Refresh  *string `json:"refresh"`
}

type UpdateUserRequest struct {
	Id        int     `json:"id" xml:"id" yaml:"id" toml:"id" query:"id" form:"id"`
	Username  string  `json:"username" xml:"username" yaml:"username" toml:"username" form:"username" query:"username"`
	Password  string  `json:"password" xml:"password" yaml:"password" toml:"password" form:"password" query:"password"`
	Role      string  `json:"role" xml:"role" yaml:"role" toml:"role" form:"role" query:"role"`
	Status    bool    `json:"status" xml:"status" yaml:"status" toml:"status" form:"status" query:"status"`
	Refresh   *string `json:"refresh" xml:"refresh" yaml:"refresh" toml:"refresh" form:"refresh" query:"refresh"`
	UpdatedBy int     `json:"-" bun:"updated_by"`
}

type UpdateUserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
	Refresh  string `json:"refresh"`
}

type UpdateUserColumnsRequest struct {
	ID     int               `json:"id" xml:"id" yaml:"id" toml:"id" form:"id" query:"id"`
	Fields map[string]string `json:"fields" xml:"fields" yaml:"fields" toml:"fields" form:"fields" query:"fields"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

type GetUserResponse struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"-"`
	Role     string  `json:"role"`
	Status   bool    `json:"status"`
	Refresh  *string `json:"refresh"`
}

type Filter struct {
	Limit int64 `json:"limit" xml:"limit" yaml:"limit" toml:"limit" form:"limit" query:"limit"`
	Page  int64 `json:"page" xml:"page" yaml:"page" toml:"page" form:"page" query:"page"`
}

type ListUserResponse struct {
	Users []*GetUserResponse `json:"users"`
	Total int                `json:"total"`
}
