package model

type Category struct {
	ID           int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaCategory string `gorm:"type:varchar(255)" json:"nama_category"`
}