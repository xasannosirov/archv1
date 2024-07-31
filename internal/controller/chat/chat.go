package chat

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/config"
	"archv1/internal/pkg/errors"
	"archv1/internal/pkg/repo/postgres"
	"archv1/internal/pkg/repo/redis"
	"archv1/internal/pkg/utils"
	"archv1/internal/usecase/chat"
	"archv1/internal/websocket"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
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

func (ch *ChatController) SendMessage(c *gin.Context) {}

func (ch *ChatController) UpdateMessage(c *gin.Context) {}

func (ch *ChatController) DeleteMessage(c *gin.Context) {}

func (ch *ChatController) GetChatMessages(c *gin.Context) {}

func (ch *ChatController) GetNotifications(c *gin.Context) {}
