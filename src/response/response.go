package response

import "github.com/danielmoisa/envoy/src/model"

type Response interface {
	ExportForFeedback() interface{}
}

type LoginResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}
