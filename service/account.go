package service

import (
	"github.com/bbcyyb/goarch/entity"
	infra "github.com/bbcyyb/goarch/infrastructure"
	repo "github.com/bbcyyb/goarch/repository"
)

type AccountService interface {
	AddOrUpdateUser(user entity.User, canSave bool) error
	ValidateEmailId(emailId string, id int) (bool, error)
	Get(emailId string) (*entity.User, error)
}

type accountDefaultService struct {
	UserRepo repo.UserRepository
}

func NewUserService() *UserService {
	return newUserDefaultService()
}

func newUserDefaultService() *userDefaultService {
	infra.C.Get(xxxxxxxxx)
	return &userDefaultService{}
}

func (self *userDefaultService) AddOrUpdateUser(user entity.User, canSave bool) error {

}

func (self *userDefaultService) ValidateEmailId(emailId string, id int) (bool, error) {

}

func (self *userDefaultService) Get(emailId string) (*entity.User, error) {

}
