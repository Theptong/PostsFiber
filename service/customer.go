package service

import "workshop/structs"

type CustomerService interface {
	GetCustomerService() ([]structs.Posts, error)
	GetCollectionService() (structs.ListPosts)
	GetCustomerServiceById(id string) (*structs.Posts, error)
}

