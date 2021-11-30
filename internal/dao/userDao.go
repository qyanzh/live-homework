package dao

import (
	"live/internal/entity"
)

type UserDao interface {
	AddUser(user *entity.User) uint32
	DeleteUser(user *entity.User) uint32
	UpdateUser(user *entity.User) uint32
	QueryUserByID(ID uint32) *entity.User
}
