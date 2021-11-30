package dao

import (
	"gorm.io/gorm"
	"live/internal/entity"
)

type UserDaoImpl struct {
	db *gorm.DB
}

func NewUserDaoImpl() UserDao {
	return &UserDaoImpl{db: DbClient}
}

func (u *UserDaoImpl) AddUser(user *entity.User) uint32 {
	u.db.Create(&user)
	return user.ID
}

func (u *UserDaoImpl) DeleteUser(user *entity.User) uint32 {
	result := u.db.Delete(&user)
	return uint32(result.RowsAffected)
}

func (u *UserDaoImpl) UpdateUser(user *entity.User) uint32 {
	result := u.db.Model(&user).Updates(user)
	return uint32(result.RowsAffected)
}

func (u *UserDaoImpl) QueryUserByID(ID uint32) *entity.User {
	user := entity.User{ID: ID}
	result := u.db.First(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}
