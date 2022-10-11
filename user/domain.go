package user

import (
	"github.com/msyamsula/messaging-api/user/service"
)

type UserDomain struct {
	Svc service.Service
}
