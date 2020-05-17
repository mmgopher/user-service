package model

import "time"

// User represents User entity
type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Surname   string    `db:"surname"`
	Gender    string    `db:"gender"`
	Age       int       `db:"age"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
}
