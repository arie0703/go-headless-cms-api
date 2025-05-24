package models

type ContentType struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}
