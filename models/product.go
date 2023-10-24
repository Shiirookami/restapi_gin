package models

type Product struct {
	id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"typer:varchar(100)" json:"NamaProduct"`
	Deskripsi   string `gorm:"type:text" json:"Deskripsi"`
}
