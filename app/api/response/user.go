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

// UserListWithPagination represents json response for GET /users route.
type UserListWithPagination struct {
	Result     []User     `json:"result"`
	Pagination Pagination `json:"pagination"`
}

// Pagination represents pagination data
type Pagination struct {
	PrevLink string `json:"prev_link"`
	NextLink string `json:"next_link"`
}
