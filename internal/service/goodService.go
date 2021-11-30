package service

import (
	"live/internal/entity"
)

type GoodService interface {
	AddGood(good *entity.Good) uint32
	DeleteGood(good *entity.Good) uint32
	UpdateGood(good *entity.Good) uint32
	QueryGoodByID(ID uint32) *entity.Good
}
