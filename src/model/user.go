package model

type User struct {
	ID       string `gorm:"privateKey" json:"id"`
	UserName string `gorm:"unique" json:"username"`
	Password string `json:"-"`
}

type UserSession struct {
	JWTToken string `json:"jwt_token"`
}
