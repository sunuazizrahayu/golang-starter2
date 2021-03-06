package repositories

import (
	"golang-starter/infrastructure/db/onlinedb"
	"golang-starter/src/users/models"
)

type UserRepository interface {
	FindAll() []models.Users
	FindByID(id uint) models.Users
	FindByEmail(email string) models.Users
}

type userRepository struct {
	DB onlinedb.Database
}

func ProvideUserRepository(DB onlinedb.Database) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) FindAll() []models.Users {
	var users []models.Users
	u.DB.Query().Find(&users)

	return users
}

func (u *userRepository) FindByID(id uint) models.Users {
	var user models.Users
	u.DB.Query().First(&user, id)

	return user
}

func (u *userRepository) FindByEmail(email string) models.Users {
	var user models.Users
	u.DB.Query().Where("email = ?", email).First(&user)

	return user
}
