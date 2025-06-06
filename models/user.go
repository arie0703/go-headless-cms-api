package models

type UserRole string

const (
	Admin  UserRole = "admin"
	Editor UserRole = "editor"
	Viewer UserRole = "viewer"
)

type User struct {
	ID           string   `gorm:"type:char(36);primaryKey" json:"id"`
	Name         string   `gorm:"type:varchar(100);not null" json:"name"`
	Email        string   `gorm:"type:varchar(255);not null;unique" json:"email"`
	PasswordHash string   `gorm:"type:text;not null" json:"-"`
	Role         UserRole `gorm:"type:enum('admin','editor','viewer');default:'editor'" json:"role"`
}
