package controllers

import "github.com/mahendrafathan/go-cake/models"

type CakeRepository interface {
	FindAll() ([]models.Cake, error)
	FindOne(id int) (models.Cake, error)
	Create(models.Cake) error
	Update(id int, cake models.Cake) error
	Delete(id int) error
}
