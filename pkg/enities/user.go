package enities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
	Role     string
}

type Login struct {
	Email    string
	Password string
}
