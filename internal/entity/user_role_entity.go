package entity

type UserRole struct {
	BaseModelUUID
	UserID    string `gorm:"type:uuid;not null;index"`
	RoleID    string `gorm:"type:uuid;not null;index"`

	User 	 User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role	 Role `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}