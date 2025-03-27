package controller

import (
	"net/http"
	"strconv"

	"github.com/danielmoisa/envoy/src/model"
	"github.com/danielmoisa/envoy/src/request"
	"github.com/gin-gonic/gin"
)

type UserDTO *model.User

// GetAllUsers retrieves all users
// @Summary Get all users
// @Description Fetch all users
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} UserDTO
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [get]
func (ctrl *Controller) GetAllUsers(c *gin.Context) {
	users, errInRetrieveUsers := ctrl.Repository.UsersRepository.RetrieveUsers()
	if errInRetrieveUsers != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_GET_FAILED, "Get users error: "+errInRetrieveUsers.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser retrieves a user by ID
// @Summary Get a user by ID
// @Description Fetches user details using the provided user ID
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Param userId path int true "User ID"
// @Success 200 {object} UserDTO
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{userId} [get]
func (ctrl *Controller) GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))

	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Get user error: "+err.Error())
		return
	}

	user, err := ctrl.Repository.UsersRepository.RetrieveByUserID(userID)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_GET_FAILED, "get user by id error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Param User body UserDTO true "User details"
// @Success 201 {object} UserDTO
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [post]
func (ctrl *Controller) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Invalid user data: "+err.Error())
		return
	}

	if req.Password == "" {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_CREATE_FAILED, "Password is required")
		return
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Role:     req.Role,
	}

	hashedPassword, err := user.HashPassword(req.Password)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_PASSWORD_INVALID, "Password hashing failed: "+err.Error())
		return
	}

	user.Password = hashedPassword

	user, err = ctrl.Repository.UsersRepository.Create(
		user.Username,
		user.Email,
		user.Password,
		user.Avatar,
	)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_CREATE_FAILED, "Create user error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates a user by ID
// @Summary Update a user by ID
// @Description Update a user by their unique user ID
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Param userId path int true "User ID"
// @Param User body UserDTO true "User details"
// @Success 200 {object} UserDTO
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{userId} [put]
func (ctrl *Controller) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Get user error: "+err.Error())
		return
	}

	user := &model.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_UPDATE_FAILED, "Create user error: "+err.Error())
		return
	}

	hashedPassword, err := user.HashPassword(user.Password)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_PASSWORD_INVALID, "Password hashing failed: "+err.Error())
		return
	}

	user, err = ctrl.Repository.UsersRepository.UpdateByID(userID, user.Username, user.Email, hashedPassword, user.Avatar)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_UPDATE_FAILED, "Update user by id error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
// @Summary Delete a user by ID
// @Description Delete a user by their unique user ID
// @Tags Users
// @Accept json
// @Produce json
// @Security Bearer
// @Param userId path int true "User ID"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{userId} [delete]
func (ctrl *Controller) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_INPUT, "Get user error: "+err.Error())
		return
	}

	// Fetch data
	err = ctrl.Repository.UsersRepository.DeleteByID(userID)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_USER_DELETE_FAILED, "Delete user by id error: "+err.Error())
		return
	}

	// Response
	c.JSON(http.StatusOK, nil)
}
