package db

import (
	"errors"
	"fmt"

	"github.com/NajmiddinAbdulhakim/iman/get-service/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
)

func ConnectDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(connDB.DB, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf(`failed to connecting migrate driver`)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:/home/najmiddin/go/src/github.com/NajmiddinAbdulhakim/iman/get-service/migrations",
		"postgers", driver,
	)
	if err != nil {
		return nil, fmt.Errorf(`failed to cannecting migrate: %v`, err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("failed to migrate: %v", err)
	}
	return connDB, nil
}
