package entity

type User struct {
	ID   uint32 `gorm:"primarykey"`
	Name string
}
