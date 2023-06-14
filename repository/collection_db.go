package repository

import (
	"workshop/structs"

	"github.com/jmoiron/sqlx"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]structs.PostsDB, error) {
	var customers []structs.PostsDB
	query := `select * from posts ORDER BY created_at ASC;`
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id string) (*structs.PostsDB, error) {
	var customers structs.PostsDB
	query := `SELECT * FROM posts where id = $1`
	err := r.db.Get(&customers, query, id)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func (r customerRepositoryDB) GetByTitle(id string) (*structs.PostsDB, error) {
	var customers structs.PostsDB
	query := `SELECT * FROM posts where title = $1`
	err := r.db.Get(&customers, query, id)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func (r customerRepositoryDB) GetByContent(id string) (*structs.PostsDB, error) {
	var customers structs.PostsDB
	query := `SELECT * FROM posts where content = $1`
	err := r.db.Get(&customers, query, id)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func (r customerRepositoryDB) GetByPublished(id string) ([]structs.PostsDB, error) {
	var customers []structs.PostsDB
	query := `select * from posts where published = $1`
	err := r.db.Select(&customers, query, id)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetByDate(id string, today string) ([]structs.PostsDB, error) {
	var customers []structs.PostsDB
	query := `SELECT * FROM posts where created_at between $1 and $2`
	err := r.db.Select(&customers, query, id, today)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) CreateNewCollection(title, content string, published bool) (*structs.PostsDB, error) {
	var customers structs.PostsDB
	query := `INSERT INTO posts (title, content, published)
	VALUES ($1, $2, $3)`
	err := r.db.Get(&customers, query, title, content, published)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func (r customerRepositoryDB) UpdateCollection(id, title, content string, published bool) (*structs.PostsDB, error) {
	var customers structs.PostsDB
	query := `update posts
	set title = $2 , content = $3 , published = $4 ,updated_at = current_timestamp
	where id = $1`
	err := r.db.Get(&customers, query, id, title, content, published)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func (r customerRepositoryDB) DeleteCollection(id string) (*structs.PostsDB, error) {
	var customers structs.PostsDB
	query := `delete from posts where id = $1`
	err := r.db.Get(&customers, query, id)
	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func (r customerRepositoryDB) LimitCollection(Offset, Limit int) ([]structs.PostsDB, error) {
	var customers []structs.PostsDB
	query := `SELECT * FROM posts OFFSET $1 LIMIT $2`
	err := r.db.Select(&customers, query, Offset, Limit)

	if err != nil {
		return nil, err
	}
	return customers, nil

}

func (r customerRepositoryDB) ViewCountCollection(ViewCount int, id string) (*structs.PostsDB, error) {
	var customers structs.PostsDB
	query := `update posts
	set view_count = $1
	where id = $2`
	err := r.db.Get(&customers, query, ViewCount, id)

	if err != nil {
		return nil, err
	}
	return &customers, nil

}
