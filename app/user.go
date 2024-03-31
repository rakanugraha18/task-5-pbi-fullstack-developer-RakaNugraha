package app

type User struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `json:"username" valid:"required"`
	Email        string `json:"email" valid:"required"`
	Password     string `json:"password" valid:"required"`
	ProfileImage string `json:"profile_image"`
}
