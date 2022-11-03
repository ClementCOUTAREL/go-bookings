package dbrepo

import (
	"database/sql"

	"github.com/ccoutarel/bookings/internal/config"
	"github.com/ccoutarel/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo creates a repository for a postgres db
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
