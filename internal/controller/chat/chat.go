package chat

import (
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/usecase/chat"
	"archv1/internal/websocket"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ChatController struct {
	Conf         *config.Config
	Hub          *websocket.Hub
	Postgres     *postgres.DB
	Redis        *redis.Redis
	Enforcer     *casbin.Enforcer
	ChatUseCaseI chat.ChatUseCaseI
}

func NewChatController(ch *ChatController) *ChatController {
	return &ChatController{
		Conf:         ch.Conf,
		Hub:          ch.Hub,
		Postgres:     ch.Postgres,
		Redis:        ch.Redis,
		Enforcer:     ch.Enforcer,
		ChatUseCaseI: ch.ChatUseCaseI,
	}
}

// groups

func (ch *ChatController) UserGroups(c *gin.Context) {}

func (ch *ChatController) GroupUsers(c *gin.Context) {}

func (ch *ChatController) GetGroup(c *gin.Context) {}

func (ch *ChatController) CreateGroup(c *gin.Context) {}

func (ch *ChatController) UpdateGroup(c *gin.Context) {}

func (ch *ChatController) UpdateGroupColumns(c *gin.Context) {}

func (ch *ChatController) DeleteGroup(c *gin.Context) {}

func (ch *ChatController) AddUserToGroup(c *gin.Context) {}

func (ch *ChatController) RemoveUserFromGroup(c *gin.Context) {}

// chat

func (ch *ChatController) UserChats(c *gin.Context) {}

func (ch *ChatController) CreateChat(c *gin.Context) {}

func (ch *ChatController) DeleteChat(c *gin.Context) {}

// messages

func (ch *ChatController) SendMessage(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		http.Error(c.Writer, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	ch.Hub.Broadcast <- body
	_, err = fmt.Fprintf(c.Writer, "Message sent to channel")
	if err != nil {
		return
	}
}

func (ch *ChatController) UpdateMessage(c *gin.Context) {}

func (ch *ChatController) DeleteMessage(c *gin.Context) {}

func (ch *ChatController) SendFile(c *gin.Context) {}

func (ch *ChatController) DeleteFile(c *gin.Context) {}

func (ch *ChatController) GetChatMessages(c *gin.Context) {}

func (ch *ChatController) GetGroupMessages(c *gin.Context) {}
