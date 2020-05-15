package dao

import "github.com/jmoiron/sqlx"

// UserRepositoryProvider provides an interface to work with database User entity
type UserRepositoryProvider interface {
}

// UserRepository represents object to work with  database User entity
type UserRepository struct {
	db *sqlx.DB
}
