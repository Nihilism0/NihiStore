package model

type OSSRecord struct {
	ID   string `gorm:"primarykey"`
	Path string `gorm:"column:path;type:varchar(100);not null"`
}
