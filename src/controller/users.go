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

	//TODO: return err for user
	if errInGetTeamID != nil {
		return
	}

	// Fetch data
	users, errInRetrieveUsers := ctrl.Storage.UsersStorage.RetrieveByTeamID(teamID)
	if errInRetrieveUsers != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_RESOURCE, "get users by team id error: "+errInRetrieveUsers.Error())
		return
	}

	// Response
	c.JSON(http.StatusOK, users)
}

func (ctrl *Controller) GetUser(c *gin.Context) {
	// Fetch userId param
	userID, err := strconv.Atoi(c.Param("userId"))

	//TODO: return err for user
	if err != nil {
		return
	}

	// Fetch data
	user, err := ctrl.Storage.UsersStorage.RetrieveByUserID(userID)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_CAN_NOT_GET_RESOURCE, "get users by team id error: "+err.Error())
		return
	}

	// Response
	c.JSON(http.StatusOK, user)
	return

}
