package model

type Product struct {
	ID               int64    `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaProduct      string   `gorm:"type:varchar(255)" json:"nama_product"`
	Harga            int64    `gorm:"type:int(10)" json:"harga"`
	GambarProduct    string   `gorm:"type:text" json:"gambar_product"`
	DeskripsiProduct string   `gorm:"type:text" json:"deskripsi_product"`
	TahunProduct     int64    `gorm:"type:int(4)" json:"tahub_product"`
	CategoryID       int64    `gorm:"index" json:"category_id"`
	Category         Category `gorm:"foreignKey:CategoryID" json:"category"`
}