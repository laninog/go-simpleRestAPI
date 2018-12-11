package repositories

import (
	"errors"
	"log"
	"strconv"

	"github.com/laninog/go-simpleRestAPI/models"
)

type usersRepository struct {
	sequence int
	users []models.User
}

func NewUserRepository() *usersRepository {
	repo := &usersRepository{
		sequence: 1,
	}
	repo.users = append(repo.users, models.User{"0", "FirstName", "LastName", &models.Address{"City", "State"}})
	return repo
}

func (r *usersRepository) add(user models.User) models.User {
	log.Printf("Add User %v", user)
	user.ID = strconv.Itoa(r.sequence)
	r.sequence++
	r.users = append(r.users, user)
	return user
}

func (r *usersRepository) remove(ID string) (models.User, error) {
	log.Printf("Remove User %v", ID)
	for index, u := range r.users {
		if u.ID == ID {
			r.users = append(r.users[:index], r.users[index+1:]...)
			return u, nil
		}
	}
	return models.User{}, errors.New("NOT FOUND")
}

func (r *usersRepository) update(ID string, user models.User) (models.User, error) {
	_, err := r.remove(ID)
	if err != nil {
		user.ID = ID
		r.add(user)
		return user, nil
	}
	return models.User{}, errors.New("NOT FOUND")
}

func (r *usersRepository) findByID(ID string) (models.User, error) {
	log.Printf("Find User by ID %v", ID)
	for _, u := range r.users {
		if u.ID == ID {
			return u, nil
		}
	}
	return models.User{}, errors.New("NOT FOUND")
}

func (r *usersRepository) findAll() []models.User {
	log.Printf("Find All Users")
	return r.users
}