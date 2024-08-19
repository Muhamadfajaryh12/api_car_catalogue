package model

type Product struct {
	ID               int64    `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	NamaProduct      string   `gorm:"type:varchar(255)" json:"nama_product" form:"nama_product"`
	Harga            int64    `gorm:"type:int(10)" json:"harga" form:"harga"`
	GambarProduct    string   `gorm:"type:text" json:"gambar_product" form:"gambar_product"`
	DeskripsiProduct string   `gorm:"type:text" json:"deskripsi_product" form:"deskripsi_product"`
	TahunProduct     int64    `gorm:"type:int(4)" json:"tahun_product" form:"tahun_product"`
	CategoryID       int64    `gorm:"index" json:"category_id" form:"category_id"`
	Category         Category `gorm:"foreignKey:CategoryID" json:"category"`
}