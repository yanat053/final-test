package entity

import(
	"gorm.io/gorm"
)
type User struct {
	gorm.Model

	Username string `json:"username" valid:"required~Username is required"`
	Email    string `json:"email" valid:"email~Invalid email format,required~Email is required"`
	Password string `json:"password" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters long"`
}