package storage

import (
	"github.com/NajmiddinAbdulhakim/iman/get-service/storage/network"
	"github.com/NajmiddinAbdulhakim/iman/get-service/storage/postgres"
	"github.com/NajmiddinAbdulhakim/iman/get-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	GetFromAPI() repo.GetAPIStorage
	SQLMethods() repo.SQLStorage
}

type storagePg struct {
	db         *sqlx.DB
	getAPI     repo.GetAPIStorage
	sqlMethods repo.SQLStorage
}

func NewStoragePg(db *sqlx.DB, link string) *storagePg {
	return &storagePg{
		db:         db,
		getAPI:     network.New(link),
		sqlMethods: postgres.New(db),
	}
}

func (s *storagePg) GetFromAPI() repo.GetAPIStorage {
	return s.getAPI
}

func (s *storagePg) SQLMethods() repo.SQLStorage {
	return s.sqlMethods
}
