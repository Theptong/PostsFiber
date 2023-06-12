package handler

import (
	"workshop/structs"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type CustomerHandler interface {
	Create(c *fiber.Ctx, db *sqlx.DB) (*structs.Posts, error)
	GetById(id string, c *fiber.Ctx, db *sqlx.DB) (*structs.PostsDB, error)
	GetByListId(id string, c *fiber.Ctx, db *sqlx.DB) (*structs.PostsDB, error)
	GetCollection(c *fiber.Ctx, db *sqlx.DB) (structs.ListPosts, error)
	UpdateById(id string, c *fiber.Ctx, db *sqlx.DB) (*structs.Posts, error)
	DeleteById(id string, c *fiber.Ctx, db *sqlx.DB)  error
}
