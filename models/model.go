package models

var (
	Viewer = "GET"
	Editor = "^(POST|PUT|PATCH)$"
	Admin  = "*"
)

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(255);not null;unique"`
	Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID          uint         `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"type:varchar(255);not null;unique"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Action  string `gorm:"type:varchar(255);not null"` // GET, POST
	Route   string `gorm:"type:varchar(255);not null"` // e.g. /api/v1/user, /api/v1/user/*
	Allowed bool   `gorm:"not null"`                   // Allow or Deny
}
