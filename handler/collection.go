package handler

import (
	"workshop/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type CustomerHandler interface {
	Create(c *fiber.Ctx, db *sqlx.DB) (*structs.Posts, error)
	// GetById(id string) (*structs.PostsDB, error)
	// GetByTitle(id string) (*structs.PostsDB, error)
	// GetByContent(id string) (*structs.PostsDB, error)
	// GetByPublished(id bool) (*structs.PostsDB, error)
}
