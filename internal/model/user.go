package model

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:255;not null;unique"`
	Email    string `gorm:"size:255;not null;unique"`
}
