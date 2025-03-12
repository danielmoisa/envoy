package controller

import (
	"net/http"
	"strconv"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/gin-gonic/gin"
)

type UserDTO *model.User

// GetAllUsers retrieves all users
// @Summary Get all users
// @Description Fetch all users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} UserDTO
// @Failure 400
// @Failure 500
// @Router /users [get]
func (ctrl *Controller) GetAllUsers(c *gin.Context) {

	// Fetch data
	users, errInRetrieveUsers := ctrl.Repository.UsersRepository.RetrieveUsers()
	if errInRetrieveUsers != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_TEAM, "get users error: "+errInRetrieveUsers.Error())
		return
	}

	// Response
	c.JSON(http.StatusOK, users)
}

// GetUser retrieves a user by ID
// @Summary Get a user by ID
// @Description Fetches user details using the provided user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} UserDTO "User details"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{userId} [get]
func (ctrl *Controller) GetUser(c *gin.Context) {
	// Fetch param
	userID, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_TEAM, "get users by team id error: "+err.Error())
		return
	}

	// Fetch data
	user, err := ctrl.Repository.UsersRepository.RetrieveByUserID(userID)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_RESOURCE, "get user by id error: "+err.Error())
		return
	}

	// Response
	c.JSON(http.StatusOK, user)
}

// CreateUser
// @Tags Users
// @Accept json
// @Produce json
// @Param User body UserDTO true "User details"
// @Success 201 {object} UserDTO "User created successfully"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [post]
// @Summary Create a new user
// @Description Create a new user
func (ctrl *Controller) CreateUser(c *gin.Context) {
	user := &model.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_RESOURCE, "create user error: "+err.Error())
		return
	}

	user, err := ctrl.Repository.UsersRepository.Create(user.Nickname, user.Email, user.PasswordDigest, user.Avatar)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_RESOURCE, "create user error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
