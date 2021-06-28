package database

import (
	"time"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

func newConfigMongoDB() *config {
	return &config{
		host:       "localhost",
		database:   "bank",
		password:   "password123",
		user:       "root",
		ctxTimeout: 60 * time.Second,
	}
}

func newConfigPostgres() *config {
	return &config{
		host:     "mord_db",
		database: "vhry",
		port:     "5432",
		driver:   "postgres",
		user:     "root",
		password: "VTwUrUEeuyBTL",
	}
}
