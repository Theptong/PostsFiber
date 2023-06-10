package database

import (
	"workshop/structs"

	"github.com/jmoiron/sqlx"
)

type collectionsRepositoryDB struct {
	db *sqlx.DB
}

func NewCollectionsRepositoryDB(db *sqlx.DB) collectionsRepositoryDB {
	return collectionsRepositoryDB{db: db}
}

func (r collectionsRepositoryDB) GetAll() ([]structs.Posts, error) {
	var collections []structs.Posts
	query := `select * from posts`
	err := r.db.Select(&collections, query)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (r collectionsRepositoryDB) GetById(id string) (*structs.PostsDB, error) {
	var customers structs.PostsDB
	query := `SELECT * FROM posts where id = $1`
	err := r.db.Get(&customers, query, id)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}