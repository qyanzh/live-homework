package dao

import (
	"gorm.io/gorm"
	"live/internal/entity"
)

type RoomDaoImpl struct {
	db *gorm.DB
}

func NewRoomDaoImpl() RoomDao {
	return &RoomDaoImpl{db: DbClient}
}

func (u *RoomDaoImpl) AddRoom(room *entity.Room) uint32 {
	u.db.Create(&room)
	return room.ID
}

func (u *RoomDaoImpl) DeleteRoom(room *entity.Room) uint32 {
	result := u.db.Delete(&room)
	return uint32(result.RowsAffected)
}

func (u *RoomDaoImpl) UpdateRoom(room *entity.Room) uint32 {
	result := u.db.Model(&room).Updates(room)
	return uint32(result.RowsAffected)
}

func (u *RoomDaoImpl) QueryRoomByID(ID uint32) *entity.Room {
	room := entity.Room{ID: ID}
	result := u.db.First(&room)
	if result.Error != nil {
		return nil
	}
	return &room
}
