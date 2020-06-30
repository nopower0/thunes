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

type IUserManager interface {
	Get(uid int) (*User, error)
	GetMany(uids []int) (map[int]*User, error)
	GetByCredential(username, password string) (*User, error)
	Create(username, password string) (*User, error)
}

type UserManager struct {
	db *xorm.Engine
}

func NewUserManager(db *xorm.Engine) *UserManager {
	return &UserManager{db: db}
}

func (m *UserManager) Get(uid int) (*User, error) {
	if users, err := m.GetMany([]int{uid}); err != nil {
		return nil, err
	} else {
		return users[uid], nil
	}
}

func (m *UserManager) GetMany(uids []int) (map[int]*User, error) {
	var users []*User
	if err := m.db.In("uid", uids).Find(&users); err != nil {
		return nil, errors.Wrap(err, "error getting user by id from DB")
	}
	result := make(map[int]*User, len(users))
	for _, u := range users {
		result[u.UID] = u
	}
	return result, nil
}

func (m *UserManager) GetByCredential(username, password string) (*User, error) {
	user := new(User)
	if exist, err := m.db.Where("username = ?", username).Get(user); err != nil {
		return nil, errors.Wrap(err, "error getting user from DB")
	} else if !exist {
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
