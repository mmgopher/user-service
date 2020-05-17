package response

import "time"

// User stores response for GET /v1/users/:user_id endpoint
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Gender    string    `json:"gender"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateUser stores response for POST /users endpoint
type CreateUser struct {
	ID int `json:"id"`
}
