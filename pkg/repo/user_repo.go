package repo

import (
	"github.com/ratheeshkumar/microservice_grpc_userV1/pkg/enities"
	"gorm.io/gorm"
)

// Userrepository for connecting db to UserREpointer method
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser create new user in db otherwise return the error
func (u *UserRepository) CreateUser(user *enities.User) error {
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FinduserByID find the user in the db  using ID
func (u *UserRepository) FindUserByID(id uint) (*enities.User, error) {
	var user *enities.User
	err := u.DB.First(&user, id).Error
	return user, err
}

// UpdateUser updates user details in db
func (u *UserRepository) UpdateUser(user *enities.User) error {
	err := u.DB.Save(&user).Error
	return err
}

// DeleteUser delete a user using the id
func (u *UserRepository) DeleteUser(id uint) error {
	err := u.DB.Delete(&enities.User{}, id).Error
	return err
}

// FindUserByEmail find the user Email
func (u *UserRepository) FindUserByEmail(email string) (*enities.User, error) {
	var user *enities.User
	err := u.DB.Where("email=?", email).First(&user).Error
	return user, err
}

// FindAllUser find the all the users details in db
func (u *UserRepository) FindAllUsers() ([]*enities.User, error) {
	var users []*enities.User
	err := u.DB.Find(&users).Error
	return users, err
}
