package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/danielmoisa/envoy/src/request"
	"github.com/danielmoisa/envoy/src/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	ERROR_FLAG_INVALID_LOGIN       = "INVALID_LOGIN_REQUEST"
	ERROR_FLAG_INVALID_CREDENTIALS = "INVALID_CREDENTIALS"
	ERROR_FLAG_TOKEN_GENERATION    = "TOKEN_GENERATION_ERROR"
)

// Login authenticates a user and returns a JWT token
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body request.LoginRequest true "Login Credentials"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /auth/login [post]
func (ctrl *Controller) Login(c *gin.Context) {
	var loginReq request.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_LOGIN, "invalid login request: "+err.Error())
		return
	}

	// Find user
	user, err := ctrl.Repository.UsersRepository.FindByEmail(loginReq.Email)
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_CREDENTIALS, "invalid credentials")
		return
	}

	// Check password
	if !user.CheckPassword(loginReq.Password) {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_INVALID_CREDENTIALS, "invalid credentials")
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hour expiry
	})

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("ENVOY_JWT_SECRET")))
	if err != nil {
		ctrl.FeedbackBadRequest(c, ERROR_FLAG_TOKEN_GENERATION, "could not generate token: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, response.LoginResponse{
		Token: tokenString,
		User:  *user,
	})
}

// Logout handles user logout
// @Summary User logout
// @Description Logout user (client-side token removal)
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Successfully logged out"
// @Router /auth/logout [post]
func (ctrl *Controller) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
