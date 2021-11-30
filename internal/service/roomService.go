package service

import (
	"live/internal/entity"
)

type RoomService interface {
	AddRoom(room *entity.Room) uint32
	DeleteRoom(room *entity.Room) uint32
	UpdateRoom(room *entity.Room) uint32
	QueryRoomByID(ID uint32) *entity.Room
	QueryTurnoverByID(ID uint32) (uint32, error)
}
