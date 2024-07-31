package entity

type Groups struct {
	ID          *int    `bun:"id"`
	Name        string  `bun:"name"`
	Username    string  `bun:"username"`
	Description *string `bun:"description"`
	CreatedBy   int     `bun:"created_by"`
}

type ResponseWithStatus struct {
	Status bool `json:"status"`
}

type ResponseWithMessage struct {
	Message string `json:"message"`
}

type GetGroupResponse struct {
	GroupId     int    `json:"group_id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Description string `json:"description"`
	CreatedBy   int    `json:"-"`
}

type CreateGroupResponse struct {
	GroupID     int    `json:"group_id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

type UpdateGroupRequest struct {
	GroupID     int    `json:"group_id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Description string `json:"description"`
	UpdatedBy   int    `json:"-"`
}

type UpdateGroupResponse struct {
	GroupID     int    `json:"group_id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

type UpdateGroupColumns struct {
	GroupID int               `json:"group_id"`
	Fields  map[string]string `json:"fields"`
}

type DeleteGroupResponse struct {
	Message string `json:"message"`
}

type CreatedChatResponse struct {
	ChatId   int    `json:"chat_id"`
	ChatType string `json:"chat_type"`
	Creator  int    `json:"creator"`
}

type UserChatsResponse struct {
	Chats []struct {
		ChatID   int    `json:"chat_id"`
		ChatType string `json:"chat_type"`
		Creator  int    `json:"creator"`
	}
}
