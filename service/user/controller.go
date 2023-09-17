package user

import "GoCloud/pkg/serializer"

type Service interface {
	Register() serializer.Response
}

func NewUserService() Service {
	return &Param{}
}
