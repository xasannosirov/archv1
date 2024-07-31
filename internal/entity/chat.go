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

type SendMessageRequest struct {
	Action   string `json:"action"`
	Property struct {
		ChatID   *int   `json:"chat_id"`
		ChatType string `json:"chat_type"`
		Message  string `json:"message"`
		Sender   int    `json:"sender"`
	} `json:"property"`
}

type UpdateMessageRequest struct {
	Action   string `json:"action"`
	Property struct {
		ChatID     int    `json:"chat_id"`
		ChatType   string `json:"chat_type"`
		MessageID  int    `json:"message_id"`
		NewMessage string `json:"new_message"`
	}
}

type DeleteMessageRequest struct {
	Action   string `json:"action"`
	Property struct {
		ChatID    int    `json:"chat_id"`
		ChatType  string `json:"chat_type"`
		MessageID int    `json:"message_id"`
	}
}

type ChatMessagesResponse struct {
	Messages []struct {
		ChatID   int    `json:"chat_id"`
		ChatType string `json:"chat_type"`
		Sender   int    `json:"sender"`
		Message  string `json:"content"`
	}
}

type NotificationsResponse struct {
	Notifications []struct {
		ChatID             int    `json:"chat_id"`
		ChatType           string `json:"chat_type"`
		LatestSender       int    `json:"latest_sender"`
		LatestMessage      string `json:"latest_message"`
		TotalMessagesCount int    `json:"total_messages_count"`
	}
}
