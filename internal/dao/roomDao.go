package dao

import (
	"live/internal/entity"
)

type RoomDao interface {
	AddRoom(room *entity.Room) uint32
	DeleteRoom(room *entity.Room) uint32
	UpdateRoom(room *entity.Room) uint32
	QueryRoomByID(ID uint32) *entity.Room
}
