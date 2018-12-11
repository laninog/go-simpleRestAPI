package repositories

import "github.com/laninog/go-simpleRestAPI/models"

type Repository interface {

	Add(user *models.User) *models.User
	Remove(ID string) (*models.User, error)
	Update(ID string, user *models.User) (*models.User, error)
	FindByID(ID string) (*models.User, error)
	FindAll() *[]models.User

}