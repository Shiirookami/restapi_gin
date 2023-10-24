package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(100)" json:"NamaProduct"`
	Deskripsi   string `gorm:"type:text" json:"Deskripsi"`
}
