package models

import "gorm.io/gorm"

var (
	Viewer = "GET"
	Editor = "^(POST|PUT|PATCH)$"
	Admin  = "*"
)

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Roles []Role `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Role struct {
	ID          uint         `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"type:varchar(255);not null;uniqueIndex"`
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Permission struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	// GET, POST
	Action string `gorm:"type:varchar(255);not null;index:idx_action_route;check:action_check"`

	// e.g. /api/v1/user, /api/v1/user/*
	Route string `gorm:"type:varchar(255);not null;index:idx_action_route;check:route_check"`

	// Allow or Deny
	Allowed bool `gorm:"not null"`
}

func (Permission) TableName() string {
	return "permissions"
}

func (Permission) CheckConstraints(db *gorm.DB) {
	db.Exec("ALTER TABLE permissions ADD CONSTRAINT action_check CHECK (Action IN ('GET', 'POST', 'PUT', 'PATCH', 'DELETE'))")
	db.Exec("ALTER TABLE permissions ADD CONSTRAINT route_check CHECK (Route LIKE '/%')")
}
