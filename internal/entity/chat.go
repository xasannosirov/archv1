package entity

// groups

type Groups struct {
	ID          *int    `bun:"id"`
	Name        string  `bun:"name"`
	Username    string  `bun:"username"`
	Description *string `bun:"description"`
	CreatedBy   int     `bun:"created_by"`
}

type ListGroupResponse struct {
	Groups     []*GetGroupResponse `json:"groups"`
	TotalCount uint64              `json:"total_count"`
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

// chat

type UserChatsResponse struct {
	Chats []struct {
		ChatID  int `json:"chat_id"`
		Creator int `json:"creator"`
	}
}

type CreateChatRequest struct {
	Creator int `json:"creator"`
}

type CreateChatResponse struct {
	ChatId  int    `json:"chat_id"`
	Creator string `json:"creator"`
}

type DeleteChatResponse struct {
	Message string `json:"message"`
}

// messages

type SendMessageRequest struct {
	ChatId      int    `json:"chat_id"`
	ChatType    string `json:"chat_type"`
	Sender      int    `json:"sender"`
	Content     string `json:"content"`
	MessageType string `json:"message_type"`
}

type SendMessageResponse struct {
	Action   string `json:"action"`
	Property struct {
		ChatID    int    `json:"chat_id"`
		ChatType  string `json:"chat_type"`
		Sender    int    `json:"sender"`
		MessageID int    `json:"message_id"`
		Content   string `json:"content"`
		SendTime  string `json:"send_time"`
	}
}

type UpdateMessageRequest struct {
	ChatId     int    `json:"chat_id"`
	ChatType   string `json:"chat_type"`
	Sender     int    `json:"sender"`
	MessageID  int    `json:"message_id"`
	NewContent string `json:"new_content"`
}

type UpdateMessageResponse struct {
	Action   string `json:"action"`
	Property struct {
		ChatID     int    `json:"chat_id"`
		ChatType   string `json:"chat_type"`
		Sender     int    `json:"sender"`
		MessageID  int    `json:"message_id"`
		NewContent string `json:"new_content"`
		SendTime   string `json:"send_time"`
	}
}

type DeleteMessageResponse struct {
	Action   string `json:"action"`
	Property struct {
		ChatID    int    `json:"chat_id"`
		ChatType  string `json:"chat_type"`
		Sender    int    `json:"sender"`
		MessageID int    `json:"message_id"`
	}
}

type SendFileMessageRequest struct {
	ChatId   int    `json:"chat_id"`
	ChatType string `json:"chat_type"`
	Sender   int    `json:"sender"`
	FileURL  string `json:"file_url"`
}

type SendFileMessageResponse struct {
	Action   string `json:"action"`
	Property struct {
		ChatID     int    `json:"chat_id"`
		ChatType   string `json:"chat_type"`
		Sender     int    `json:"sender"`
		MessageID  int    `json:"message_id"`
		NewFileUrl string `json:"new_file_url"`
		SendTime   string `json:"send_time"`
	}
}

type DeleteFileMessageResponse struct {
	Action   string `json:"action"`
	Property struct {
		ChatID    int    `json:"chat_id"`
		ChatType  string `json:"chat_type"`
		Sender    int    `json:"sender"`
		MessageID int    `json:"message_id"`
	}
}

type Message struct {
	MessageID   int    `json:"message_id"`
	Content     string `json:"content"`
	MessageType string `json:"message_type"`
	Sender      int    `json:"sender"`
	SendTime    string `json:"send_time"`
}

type ChatMessagesResponse struct {
	Messages []Message `json:"messages"`
}

type GroupMessagesResponse struct {
	Messages []Message `json:"messages"`
}
