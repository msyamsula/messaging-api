package handler

import (
	svc "github.com/msyamsula/messaging-api/user/service"
)

type Handler struct {
	userService svc.Service
}

func NewHandler(service *svc.Service) *Handler {
	handler := &Handler{
		userService: *service,
	}

	return handler
}
