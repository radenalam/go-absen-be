package entity

type RolePermission struct {
	BaseModelUUID
	RoleID      string `gorm:"type:uuid;not null;index"`
	PermissionID string `gorm:"type:uuid;not null;index"`

	Permission Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
	Role	   Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}