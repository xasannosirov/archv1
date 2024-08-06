package chat

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/cache"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/utils"
	"archv1/internal/usecase/chat"
	"archv1/internal/usecase/user"
	"archv1/internal/websocket"
	"context"
	"encoding/json"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
	"strings"
)

type ChatController struct {
	RedisCache   *cache.Redis
	Conf         *config.Config
	Hub          *websocket.Hub
	Postgres     *postgres.DB
	Enforcer     *casbin.Enforcer
	ChatUseCaseI chat.ChatUseCaseI
	UserUseCase  user.UserUseCaseI
}

func NewChatController(ch *ChatController) *ChatController {
	return &ChatController{
		RedisCache:   ch.RedisCache,
		Conf:         ch.Conf,
		Hub:          ch.Hub,
		Postgres:     ch.Postgres,
		Enforcer:     ch.Enforcer,
		ChatUseCaseI: ch.ChatUseCaseI,
		UserUseCase:  ch.UserUseCase,
	}
}

// UserGroups
// @Security 		BearerAuth
// @Summary 		User Groups
// @Description 	This API for getting groups with user id
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "User ID"
// @Success 		200 {object} []entity.GetGroupResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Router 			/v1/group/user-groups/{id} [GET]
func (ch *ChatController) UserGroups(c *gin.Context) {
	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	groups, err := ch.ChatUseCaseI.UserGroups(context.Background(), int64(userID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetGroup
// @Security 		BearerAuth
// @Summary 		Get Group
// @Description 	This API for getting group with id
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Group ID"
// @Success 		200 {object} entity.GetGroupResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Router 			/v1/group/{id} [GET]
func (ch *ChatController) GetGroup(c *gin.Context) {
	id := c.Param("id")

	groupID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	group, err := ch.ChatUseCaseI.GetGroup(context.Background(), int64(groupID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, group)
}

// CreateGroup
// @Security 		BearerAuth
// @Summary 		Create Group
// @Description 	This API for creating a new group
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.CreateGroupRequest true "Create Group Model"
// @Success 		201 {object} entity.CreateGroupResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/group [POST]
func (ch *ChatController) CreateGroup(c *gin.Context) {
	var request entity.CreateGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, ch.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	request.CreatedBy = cast.ToInt(claims["sub"])

	response, err := ch.ChatUseCaseI.CreateGroup(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, response)
}

// UpdateGroup
// @Security 		BearerAuth
// @Summary 		Update Group
// @Description 	This API for updating a group
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.UpdateGroupRequest true "Update Group Model"
// @Success 		200 {object} entity.UpdateGroupResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/group [PUT]
func (ch *ChatController) UpdateGroup(c *gin.Context) {
	var request entity.UpdateGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, ch.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	request.UpdatedBy = cast.ToInt(claims["sub"])

	response, err := ch.ChatUseCaseI.UpdateGroup(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateGroupColumns
// @Security 		BearerAuth
// @Summary 		Update Group Columns
// @Description 	This API for updating a group columns
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			request body entity.UpdateGroupColumns true "Update Group Columns Model"
// @Success 		200 {object} entity.UpdateGroupResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/group [PATCH]
func (ch *ChatController) UpdateGroupColumns(c *gin.Context) {
	var request entity.UpdateGroupColumns

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, ch.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	request.Fields["updated_by"] = cast.ToString(claims["sub"])

	response, err := ch.ChatUseCaseI.UpdateGroupColumns(context.Background(), request)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteGroup
// @Security 		BearerAuth
// @Summary 		Delete Group
// @Description 	This API for deleting a group
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Success 		200 {object} entity.DeleteGroupResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/group/{id} [DELETE]
func (ch *ChatController) DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	groupID, err := strconv.Atoi(id)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	claims, err := utils.GetTokenClaimsFromHeader(c.Request, ch.Conf)
	if err != nil {
		errors.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	deletedBy := cast.ToInt(claims["sub"])

	response, err := ch.ChatUseCaseI.DeleteGroup(context.Background(), int64(groupID), int64(deletedBy))
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// AddUserToGroup
// @Security 		BearerAuth
// @Summary 		Add User to Group
// @Description 	This API for adding a new user to group
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			user_id query int true "User ID"
// @Param 			group_id query int true "Group ID"
// @Success 		201 {object} entity.ResponseWithStatus
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/group/add-user [POST]
func (ch *ChatController) AddUserToGroup(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	groupID, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ch.ChatUseCaseI.AddUserToGroup(context.Background(), int64(userID), int64(groupID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.ResponseWithStatus{
		Status: true,
	})
}

// RemoveUserFromGroup
// @Security 		BearerAuth
// @Summary 		Remove User from Group
// @Description 	This API for removing user from group
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			user_id query int true "User ID"
// @Param 			group_id query int true "Group ID"
// @Success 		201 {object} entity.ResponseWithStatus
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/group/remove-user [DELETE]
func (ch *ChatController) RemoveUserFromGroup(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	groupID, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ch.ChatUseCaseI.RemoveUserFromGroup(context.Background(), int64(userID), int64(groupID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.ResponseWithStatus{
		Status: true,
	})
}

// UserChats
// @Security 		BearerAuth
// @Summary 		User Chats
// @Description 	This API for getting user chats
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Success 		200 {object} entity.UserChatsResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Router 			/v1/group/user-chats [GET]
func (ch *ChatController) UserChats(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	chats, err := ch.ChatUseCaseI.UserChats(context.Background(), int64(userID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, chats)
}

// DeleteChat
// @Security 		BearerAuth
// @Summary 		Delete Chat
// @Description 	This API for deleting chat with id
// @Tags			chat
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Chat ID"
// @Success 		200 {object} entity.ResponseWithMessage
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Router 			/v1/group/delete-chat [DELETE]
func (ch *ChatController) DeleteChat(c *gin.Context) {
	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = ch.ChatUseCaseI.DeleteChat(context.Background(), int64(chatID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.ResponseWithMessage{
		Message: "success",
	})
}

// SendMessage
// @Security		BearerAuth
// @Summary 		Send Message
// @Description 	This API for sending a message to chat
// @Tags 			chat
// @Accept 			json
// @Produce 		json
// @Param 			send body entity.SendMessageRequest true "Send Message Model"
// @Success 		200 {object} entity.ResponseWithStatus
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/send-message [POST]
func (ch *ChatController) SendMessage(c *gin.Context) {
	var message entity.SendMessageRequest

	if err := c.ShouldBindJSON(&message); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	message.ChatType = strings.ToLower(message.ChatType)
	if message.ChatType != "private" && message.ChatType != "group" {
		errors.ErrorResponse(c, http.StatusBadRequest, "property chat type must be 'private' or 'group'")
		return
	}

	if message.ChatID == 0 {
		newChat, err := ch.ChatUseCaseI.CreateChat(
			context.Background(),
			int64(message.Receiver),
			int64(message.Sender),
			message.ChatType,
		)

		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		message.ChatID = newChat.ChatId
	}

	if err := ch.ChatUseCaseI.SendMessage(context.Background(), message); err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if message.ChatType == "private" {
		userResponse, err := ch.UserUseCase.GetByID(context.Background(), message.Receiver)
		if err != nil {
			errors.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		var cacheData entity.NotificationsResponse

		cacheNotification, err := ch.RedisCache.Get(context.Background(), userResponse.Username)
		if err != nil {
			cacheData = entity.NotificationsResponse{}
		} else {
			err = json.Unmarshal([]byte(cacheNotification), &cacheData)
			if err != nil {
				errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}

		var isUpdate = false
		for _, notification := range cacheData.Notifications {
			if notification.ChatID == message.ChatID {
				notification.LatestMessage = message.Message
				notification.LatestSender = message.Sender
				notification.TotalMessagesCount = notification.TotalMessagesCount + 1
				isUpdate = true
				break
			}
		}

		if isUpdate == false {
			cacheData.Notifications = append(cacheData.Notifications, struct {
				ChatID             int    `json:"chat_id"`
				ChatType           string `json:"chat_type"`
				LatestSender       int    `json:"latest_sender"`
				LatestMessage      string `json:"latest_message"`
				TotalMessagesCount int    `json:"total_messages_count"`
			}{
				ChatID:             message.ChatID,
				ChatType:           message.ChatType,
				LatestSender:       message.Sender,
				LatestMessage:      message.Message,
				TotalMessagesCount: 1,
			})
		}

		sendingData, err := json.Marshal(&entity.SendMessageData{
			Action: "new_message",
			Property: struct {
				ChatID      int    `json:"chat_id"`
				ChatType    string `json:"chat_type"`
				Message     string `json:"message"`
				MessageType string `json:"message_type"`
				Sender      int    `json:"sender"`
				Receiver    int    `json:"receiver"`
			}{
				ChatID:      message.ChatID,
				ChatType:    message.ChatType,
				Message:     message.Message,
				MessageType: message.MessageType,
				Sender:      message.Sender,
				Receiver:    message.Receiver,
			},
		})
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		isSend := false
		for connection := range ch.Hub.Connections {
			if connection.Username == userResponse.Username {
				connection.Send <- sendingData
				isSend = true
				break
			}
		}

		if isSend == false {
			err := ch.RedisCache.Set(context.Background(), userResponse.Username, cacheData, 0)
			if err != nil {
				errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
	} else {
		groupResponse, err := ch.ChatUseCaseI.GetGroup(context.Background(), int64(message.Receiver))
		if err != nil {
			errors.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		groupUsers, err := ch.ChatUseCaseI.GroupUsers(context.Background(), int64(groupResponse.GroupId))
		if err != nil {
			errors.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		for _, gettingUser := range groupUsers {
			var cacheData entity.NotificationsResponse

			cacheNotification, err := ch.RedisCache.Get(context.Background(), gettingUser.Username)
			if err != nil {
				cacheData = entity.NotificationsResponse{}
			} else {
				err = json.Unmarshal([]byte(cacheNotification), &cacheData)
				if err != nil {
					errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
					return
				}
			}

			var isUpdate = false
			for _, notification := range cacheData.Notifications {
				if notification.ChatID == message.ChatID {
					notification.LatestMessage = message.Message
					notification.LatestSender = message.Sender
					notification.TotalMessagesCount = notification.TotalMessagesCount + 1
					isUpdate = true
					break
				}
			}

			if isUpdate == false {
				cacheData.Notifications = append(cacheData.Notifications, struct {
					ChatID             int    `json:"chat_id"`
					ChatType           string `json:"chat_type"`
					LatestSender       int    `json:"latest_sender"`
					LatestMessage      string `json:"latest_message"`
					TotalMessagesCount int    `json:"total_messages_count"`
				}{
					ChatID:             message.ChatID,
					ChatType:           message.ChatType,
					LatestSender:       message.Sender,
					LatestMessage:      message.Message,
					TotalMessagesCount: 1,
				})
			}

			sendingData, err := json.Marshal(&entity.SendMessageData{
				Action: "new_message",
				Property: struct {
					ChatID      int    `json:"chat_id"`
					ChatType    string `json:"chat_type"`
					Message     string `json:"message"`
					MessageType string `json:"message_type"`
					Sender      int    `json:"sender"`
					Receiver    int    `json:"receiver"`
				}{
					ChatID:      message.ChatID,
					ChatType:    message.ChatType,
					Message:     message.Message,
					MessageType: message.MessageType,
					Sender:      message.Sender,
					Receiver:    message.Receiver,
				},
			})
			if err != nil {
				errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			isSend := false
			for connection := range ch.Hub.Connections {
				if gettingUser.Username == connection.Username {
					connection.Send <- sendingData
					isSend = true
					break
				}
			}

			if isSend == false {
				err := ch.RedisCache.Set(context.Background(), gettingUser.Username, cacheData, 0)
				if err != nil {
					errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
					return
				}
			}
		}
	}

	c.JSON(http.StatusOK, entity.ResponseWithStatus{
		Status: true,
	})
}

// UpdateMessage
// @Security		BearerAuth
// @Summary 		Update Message
// @Description 	This API for updating a message in chat
// @Tags 			chat
// @Accept 			json
// @Produce 		json
// @Param 			send body entity.UpdateMessageRequest true "Update Message Model"
// @Success 		200 {object} entity.ResponseWithStatus
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/update-message [PUT]
func (ch *ChatController) UpdateMessage(c *gin.Context) {
	var request entity.UpdateMessageRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	chatResponse, err := ch.ChatUseCaseI.GetChat(context.Background(), int64(request.ChatID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	messageResponse, err := ch.ChatUseCaseI.GetMessage(context.Background(), int64(request.MessageID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	if err := ch.ChatUseCaseI.UpdateMessage(context.Background(), request); err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if chatResponse.ChatType == "private" {
		userResponse, err := ch.UserUseCase.GetByID(context.Background(), chatResponse.ReceiverID)
		if err != nil {
			errors.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		var cacheData entity.NotificationsResponse

		cacheNotification, err := ch.RedisCache.Get(context.Background(), userResponse.Username)
		if err != nil {
			cacheData = entity.NotificationsResponse{}
		} else {
			err = json.Unmarshal([]byte(cacheNotification), &cacheData)
			if err != nil {
				errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}

		var updatedCacheNotification entity.NotificationsResponse
		for _, notification := range cacheData.Notifications {
			if notification.ChatID == request.ChatID {
				notification.LatestMessage = request.NewMessage
			}
			updatedCacheNotification.Notifications = append(updatedCacheNotification.Notifications, notification)
		}

		sendingData, err := json.Marshal(&entity.SendMessageData{
			Action: "update_message",
			Property: struct {
				ChatID      int    `json:"chat_id"`
				ChatType    string `json:"chat_type"`
				Message     string `json:"message"`
				MessageType string `json:"message_type"`
				Sender      int    `json:"sender"`
				Receiver    int    `json:"receiver"`
			}{
				ChatID:      request.ChatID,
				ChatType:    chatResponse.ChatType,
				Message:     request.NewMessage,
				MessageType: messageResponse.MessageType,
				Sender:      request.Sender,
				Receiver:    chatResponse.ReceiverID,
			},
		})
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		isSend := false
		for connection := range ch.Hub.Connections {
			if connection.Username == userResponse.Username {
				connection.Send <- sendingData
				isSend = true
				break
			}
		}

		if isSend == false {
			err := ch.RedisCache.Set(context.Background(), userResponse.Username, updatedCacheNotification, 0)
			if err != nil {
				errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
	} else {
		groupResponse, err := ch.ChatUseCaseI.GetGroup(context.Background(), int64(chatResponse.ReceiverID))
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		groupUsers, err := ch.ChatUseCaseI.GroupUsers(context.Background(), int64(groupResponse.GroupId))
		if err != nil {
			errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		for _, gettingUser := range groupUsers {
			var cacheData entity.NotificationsResponse

			cacheNotification, err := ch.RedisCache.Get(context.Background(), gettingUser.Username)
			if err != nil {
				cacheData = entity.NotificationsResponse{}
			} else {
				err = json.Unmarshal([]byte(cacheNotification), &cacheData)
				if err != nil {
					errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
					return
				}
			}

			var updatedCacheNotification entity.NotificationsResponse
			for _, notification := range cacheData.Notifications {
				if notification.ChatID == request.ChatID {
					notification.LatestMessage = request.NewMessage
				}
				updatedCacheNotification.Notifications = append(updatedCacheNotification.Notifications, notification)
			}

			sendingData, err := json.Marshal(&entity.SendMessageData{
				Action: "update_message",
				Property: struct {
					ChatID      int    `json:"chat_id"`
					ChatType    string `json:"chat_type"`
					Message     string `json:"message"`
					MessageType string `json:"message_type"`
					Sender      int    `json:"sender"`
					Receiver    int    `json:"receiver"`
				}{
					ChatID:      request.ChatID,
					ChatType:    chatResponse.ChatType,
					Message:     request.NewMessage,
					MessageType: messageResponse.MessageType,
					Sender:      request.Sender,
					Receiver:    chatResponse.ReceiverID,
				},
			})
			if err != nil {
				errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}

			isSend := false
			for connection := range ch.Hub.Connections {
				if gettingUser.Username == connection.Username {
					connection.Send <- sendingData
					isSend = true
					break
				}
			}

			if isSend == false {
				err := ch.RedisCache.Set(context.Background(), gettingUser.Username, updatedCacheNotification, 0)
				if err != nil {
					errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
					return
				}
			}
		}
	}

	c.JSON(http.StatusOK, entity.ResponseWithStatus{
		Status: true,
	})
}

// DeleteMessage
// @Security		BearerAuth
// @Summary 		Delete Message
// @Description 	This API for deleting a message in chat
// @Tags 			chat
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Message ID"
// @Success 		200 {object} entity.ResponseWithStatus
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/delete-message/{id} [DELETE]
func (ch *ChatController) DeleteMessage(c *gin.Context) {}

// GetChatMessages
// @Security		BearerAuth
// @Summary 		Get Chat Messages
// @Description 	This API for getting chat messages with chat_id
// @Tags 			chat
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Chat ID"
// @Success 		200 {object} entity.ChatMessagesResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/chat-messages/{id} [GET]
func (ch *ChatController) GetChatMessages(c *gin.Context) {
	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	messages, err := ch.ChatUseCaseI.GetChatMessages(context.Background(), int64(chatID))
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, messages)
}

// GetAllNotifications
// @Security		BearerAuth
// @Summary 		Get All Notifications
// @Description 	This API for getting all chats notifications for one user
// @Tags 			chat
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Success 		200 {object} entity.NotificationsResponse
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/get-notifications/{id} [GET]
func (ch *ChatController) GetAllNotifications(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userData, err := ch.UserUseCase.GetByID(context.Background(), userID)
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	dataJson, err := ch.RedisCache.Get(context.Background(), userData.Username)
	if err != nil {
		errors.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	var response entity.NotificationsResponse

	err = json.Unmarshal([]byte(dataJson), &response)
	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteChatNotifications
// @Security		BearerAuth
// @Summary 		Delete Chat Notifications
// @Description 	This API for deleting chats notifications for one user
// @Tags 			chat
// @Accept 			json
// @Produce 		json
// @Param 			chat_id query int true "Chat ID"
// @Param 			user_id query int true "User ID"
// @Param 			count query int true "Read Message Count"
// @Success 		200 {object} entity.ResponseWithStatus
// @Failure 		400 {object} errors.Error
// @Failure 		401 {object} errors.Error
// @Failure 		403 {object} errors.Error
// @Failure 		404 {object} errors.Error
// @Failure 		500 {object} errors.Error
// @Router 			/v1/delete-chat-notifications [DELETE]
func (ch *ChatController) DeleteChatNotifications(c *gin.Context) {}
