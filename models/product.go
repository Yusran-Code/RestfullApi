package models

type Product struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:100;not null"`
	Hobi string `gorm:"size:100;not null"`
}
