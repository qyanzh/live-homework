package entity

type Room struct {
	ID   uint32 `gorm:"primarykey"`
	Name string
}
