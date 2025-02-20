package controller

import (
	"net/http"
	"strconv"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/gin-gonic/gin"
)

type UserResponse *model.User

// GetAllUsers retrieves all users for a given team
// @Summary Get all users by team ID
// @Description Fetch all users belonging to a specific team
// @Tags Users
// @Accept  json
// @Produce  json
// @Param team_id path int true "Team ID"
// @Success 200 {array} UserResponse
// @Failure 400
// @Failure 500
// @Router /users/{team_id} [get]
func (ctrl *Controller) GetAllUsers(c *gin.Context) {
	// Fetch teamID param
	teamID, errInGetTeamID := strconv.Atoi(c.Param("teamId"))

	if errInGetTeamID != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_RESOURCE, "get users by team id error: "+errInGetTeamID.Error())
		return
	}

	// Fetch data
	users, errInRetrieveUsers := ctrl.Storage.UsersStorage.RetrieveByTeamID(teamID)
	if errInRetrieveUsers != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_TEAM, "get users by team id error: "+errInRetrieveUsers.Error())
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
// @Success 200 {object} UserResponse "User details"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{userId} [get]
func (ctrl *Controller) GetUser(c *gin.Context) {
	// Fetch params
	teamID, err := strconv.Atoi(c.Param("teamId"))
	userID, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_TEAM, "get users by team id error: "+err.Error())
		return
	}

	// Fetch data
	user, err := ctrl.Storage.UsersStorage.RetrieveByUserID(teamID, userID)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_RESOURCE, "get user by id error: "+err.Error())
		return
	}

	// Response
	c.JSON(http.StatusOK, user)
}

// @Tags Users
// @Accept json
// @Produce json
// @Param User body models.User true "User details"
// @Success 201 {object} UserResponse "User created successfully"
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

	user, err := ctrl.Storage.UsersStorage.Create(user.Nickname, user.Email, user.PasswordDigest, user.Avatar)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_CREATE_RESOURCE, "create user error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
