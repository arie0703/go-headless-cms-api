package models

type Field struct {
	ID            string `gorm:"primaryKey"`
	ContentTypeID string
	Name          string
	Label         string
	Type          string
	IsRequired    bool
	Order         int
}
