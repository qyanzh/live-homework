package service

import (
	"live/internal/entity"
)

type UserService interface {
	AddUser(user *entity.User) uint32
	DeleteUser(user *entity.User) uint32
	UpdateUser(user *entity.User) uint32
	QueryUserByID(ID uint32) *entity.User
}
