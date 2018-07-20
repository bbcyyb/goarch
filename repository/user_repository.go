package repository

import (
	"github.com/bbcyyb/goarch/entity"
)

type UserRepository interface {
	AddOrUpdate(user entity.User, canSave bool) error
	ValidateEmailId(emailId string, id int) (bool, error)
	Get(emailId string) (*entity.User, error)
}

type userMysqlRepository struct {
}

func NewUserRepository() *UserRepository {
	return &userMysqlRepository{}
}

func (u *userMysqlRepository) AddOrUpdate(user entity.User, canSave bool) error {

}

func (u *userMysqlRepository) ValidateEmailId(emailId string, id int) (bool, error) {

}

func (u *userMysqlRepository) Get(emailId string) (*entity.User, error) {

}
