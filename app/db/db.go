package db

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Supported database types
const (
	PostgresType = "postgres"
)

// GetConnection returns connection based on provided database type.
// Now only Postgress is supported
func GetConnection(dbType, user, password, name, host string) (*sqlx.DB, error) {

	var err error
	var connection *sqlx.DB
	switch dbType {
	case PostgresType:
		connection, err = getPostgresConnection(user, password, name, host)
	default:
		return nil, errors.New("unsupported database type")
	}

	if err != nil {
		return nil, errors.Wrapf(err, "impossible to get %s database connection", dbType)
	}

	if err := connection.Ping(); err != nil {
		return nil, errors.Wrap(err, "impossible to reach database")
	}

	return connection, nil
}

func getPostgresConnection(user, password, name, host string) (*sqlx.DB, error) {

	return sqlx.Connect(PostgresType, fmt.Sprintf(
		`postgres://%s:%s@%s/%s?sslmode=disable`,
		user,
		url.QueryEscape(password),
		host,
		name,
	))
}
