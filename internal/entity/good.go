package entity

type Good struct {
	ID   uint32 `gorm:"primarykey"`
	Name string
}
