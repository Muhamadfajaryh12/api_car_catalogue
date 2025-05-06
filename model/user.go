package model

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" `
	Username string `gorm:"type:varchar(255)" form:"username" json:"username"`
	Password string `gorm:"type:varchar(255)" form:"password" json:"password"`
}