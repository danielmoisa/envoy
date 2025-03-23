package request

import "github.com/danielmoisa/envoy/src/model"

const (
	ACTION_REQUEST_CONTENT_FIELD_VIRTUAL_RESOURCE = "virtualResource"
	ORDER_BY_CREATED_AT                           = "createdAt"
	ORDER_BY_UPDATED_AT                           = "updatedAt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Username string         `json:"username" binding:"required,min=3,max=50"`
	Password string         `json:"password" binding:"required,min=6"`
	Email    string         `json:"email" binding:"required,email"`
	Avatar   string         `json:"avatar"`
	Role     model.UserRole `json:"role" binding:"required,oneof=candidate company admin"`
}
