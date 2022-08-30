package postgres

import (
	pb "github.com/NajmiddinAbdulhakim/iman/get-service/genproto"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) CreatePost(posts []*pb.Post) error {
	query := `INSERT INTO posts VALUES($1, $2, $3, $4)`
	for _, p := range posts {
		_, err := r.db.Exec(query, p.Id, p.UserId, p.Title, p.Body)
		if err != nil {
			return err
		}
	}
	return nil
}
