package dbrepo

import (
	"database/sql"

	"github.com/oseintow/bookings/internal/config"
	"github.com/oseintow/bookings/internal/repository"
)

type postresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postresDBRepo{
		App: a,
		DB:  conn,
	}
}
