package service

import (
	"github.com/bbcyyb/goarch/entity"
)

type UserService interface {
	AddOrUpdate(user entity.User, canSave bool) error
	ValidateEmailId(emailId string, id int) (bool, error)
	Get(emailId string) (*entity.User, error)
}

type userDefaultService struct {
}

func NewUserService() *UserService {
	return &userDefaultService{}
}

func (u *userDefaultService) AddOrUpdate(user entity.User, canSave bool) error {

}

func (u *userDefaultService) ValidateEmailId(emailId string, id int) (bool, error) {

}

func (u *userDefaultService) Get(emailId string) (*entity.User, error) {

}
