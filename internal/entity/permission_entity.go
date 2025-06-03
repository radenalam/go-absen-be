package entity

type Permission struct {
	BaseModelUUID
	Name string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Description string `gorm:"type:varchar(255);not null"`

	RolePermissions []RolePermission `gorm:"foreignKey:PermissionID"`
}