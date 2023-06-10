package handler

import "github.com/gofiber/fiber/v2"

type CustomerHandler interface {
	GetCustomers() (c *fiber.Ctx)
	// GetById(id string) (*structs.PostsDB, error)
	// GetByTitle(id string) (*structs.PostsDB, error)
	// GetByContent(id string) (*structs.PostsDB, error)
	// GetByPublished(id bool) (*structs.PostsDB, error)
}
