package entity

import "time"

type Tx struct {
	TxID     uint32 `gorm:"unique"`
	RoomID   uint32
	GoodID   uint32
	Turnover uint32
	TxTime   time.Time
	UserID   uint32
}
