package entity

type User struct {
    BaseModelUUID
    Name      string         `gorm:"type:varchar(255);not null"`
    Username  string         `gorm:"type:varchar(255);not null;uniqueIndex"`
    Password  string         `gorm:"type:varchar(255);not null"`
    FcmToken string          `gorm:"type:varchar(255)"`
    Email     *string        `gorm:"type:varchar(255);uniqueIndex"`

    RefreshTokens []RefreshToken `gorm:"foreignKey:UserID"`
    UserRoles []UserRole         `gorm:"foreignKey:UserID"`
}