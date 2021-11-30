package dao

import (
	"gorm.io/gorm"
	"live/internal/entity"
)

type GoodDaoImpl struct {
	db *gorm.DB
}

func NewGoodDaoImpl() GoodDao {
	return &GoodDaoImpl{db: DbClient}
}

func (u *GoodDaoImpl) AddGood(good *entity.Good) uint32 {
	u.db.Create(&good)
	return good.ID
}

func (u *GoodDaoImpl) DeleteGood(good *entity.Good) uint32 {
	result := u.db.Delete(&good)
	return uint32(result.RowsAffected)
}

func (u *GoodDaoImpl) UpdateGood(good *entity.Good) uint32 {
	result := u.db.Model(&good).Updates(good)
	return uint32(result.RowsAffected)
}

func (u *GoodDaoImpl) QueryGoodByID(ID uint32) *entity.Good {
	Good := entity.Good{ID: ID}
	result := u.db.First(&Good)
	if result.Error != nil {
		return nil
	}
	return &Good
}
