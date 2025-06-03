package entity

type Role struct {
	BaseModelUUID
	Name string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Description string `gorm:"type:varchar(255);not null"`

	UserRoles []UserRole `gorm:"foreignKey:RoleID"`
	RolePermissions []RolePermission `gorm:"foreignKey:RoleID"`
}