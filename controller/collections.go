package controller

import "workshop/structs"

type CustomerRepository interface {
	GetAll() ([]structs.PostsDB, error)
	GetById(id int) (*structs.PostsDB, error)
}