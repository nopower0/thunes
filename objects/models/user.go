package models

import (
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

type User struct {
	AbstractTimeModel `xorm:"extends"`
	UID               int    `xorm:"bigint pk autoincr 'uid'" json:"uid"`
	Username          string `xorm:"varchar(100)" json:"username"`
	Password          string `xorm:"varchar(64)" json:"password"`
}

type UserManager struct {
	db *xorm.Engine
}

func NewUserManager(db *xorm.Engine) *UserManager {
	return &UserManager{db: db}
}

func (m *UserManager) Get(username, password string) (*User, error) {
	user := new(User)
	if exist, err := m.db.Where("username = ?", username).Get(user); err != nil {
		return nil, errors.Wrap(err, "error getting user from DB")
	} else  if !exist {
		return nil, nil
	}
	if user.Password != password {
		return nil, nil
	}
	return user, nil
}

func (m *UserManager) Create(username, password string) (*User, error) {
	user := &User{
		Username: username,
		Password: password,
	}
	if _, err := m.db.InsertOne(user); err != nil {
		return nil, err
	}
	return user, nil
}
