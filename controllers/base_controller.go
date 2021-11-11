package controllers

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Controller struct {
		DB *pgxpool.Pool
	}
)