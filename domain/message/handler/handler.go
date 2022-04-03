package handler

import "github.com/msyamsula/messaging-api/domain/message/service"

type Handler struct {
	messageSvc *service.Service
}

func NewHandler(messageSvc *service.Service) *Handler {
	svc := &Handler{
		messageSvc: messageSvc,
	}

	return svc
}
