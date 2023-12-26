package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]User, error)
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
