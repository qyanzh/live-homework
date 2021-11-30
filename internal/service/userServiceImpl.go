package service

import (
	"live/internal/dao"
	"live/internal/entity"
)

type UserServiceImpl struct {
	userDao dao.UserDao
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{userDao: dao.NewUserDaoImpl()}
}

func (u *UserServiceImpl) AddUser(user *entity.User) uint32 {
	return u.userDao.AddUser(user)
}

func (u *UserServiceImpl) DeleteUser(user *entity.User) uint32 {
	return u.userDao.DeleteUser(user)
}

func (u *UserServiceImpl) UpdateUser(user *entity.User) uint32 {
	return u.userDao.UpdateUser(user)
}

func (u *UserServiceImpl) QueryUserByID(ID uint32) *entity.User {
	return u.userDao.QueryUserByID(ID)
}
