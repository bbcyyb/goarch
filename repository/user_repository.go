package repository

import (
	"github.com/bbcyyb/goarch/entity"
	infra "github.com/bbcyyb/goarch/infrastructure"
)

type UserRepository interface {
	AddOrUpdate(user entity.User, canSave bool) error
	ValidateEmailId(emailId string, id int) (bool, error)
	Get(emailId string) (*entity.User, error)
}

type userMysqlRepository struct {
	conn infra.DB
}

func NewUserRepository() *UserRepository {
	return &userMysqlRepository{conn: infra.C.DbConn}
}

func (self *userMysqlRepository) AddOrUpdate(user entity.User, canSave bool) error {

}

func (self *userMysqlRepository) ValidateEmailId(emailId string, id int) (bool, error) {

}

func (self *userMysqlRepository) Get(emailId string) (*entity.User, error) {

}
