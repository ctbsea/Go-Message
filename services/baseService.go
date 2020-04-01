package services

import (
	"github.com/ctbsea/Go-Message/repositories"
)

type Service struct {
	UserService UserService
}

func InitService(rep *repositories.Rep) *Service {
	return &Service{
		UserService: NewUserService(rep),
	}
}
