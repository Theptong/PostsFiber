package controller

import (
	"workshop/structs"
	"workshop/database"
	"github.com/jmoiron/sqlx"
)

type collectionsRepositoryDB struct {
	db *sqlx.DB
}

func NewCollectionsRepositoryDB(db *sqlx.DB) collectionsRepositoryDB {
	return collectionsRepositoryDB{db: db}
}

// func (r collectionsRepositoryDB) GetAll() ([]structs.Posts, error) {
// 	var collections []structs.Posts
// 	query := `select * from posts`
// 	err := r.db.Select(&collections, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return collections, nil
// }

// GET
func (r collectionsRepositoryDB) GetCollection(db *sqlx.DB) structs.ListPosts {
	var dataList structs.ListPosts
	data := database.NewCollectionsRepositoryDB(db)
	rows, _ := data.GetAll()

	dataList.Posts = rows
	dataList.Count = len(rows)
	dataList.Limit = len(rows)
	total := (dataList.Count / dataList.Limit)
	dataList.Page = total
	remainder := (dataList.Count % dataList.Limit)
	if remainder == 0 {
		dataList.TotalPage = total
	} else {
		dataList.TotalPage = total + 1
	}

	return dataList
}

// func (r collectionsRepositoryDB) GetById(id string) (*structs.PostsDB, error) {
// 	var customers structs.PostsDB
// 	query := `SELECT * FROM posts where id = $1`
// 	err := r.db.Get(&customers, query, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &customers, nil
// }
