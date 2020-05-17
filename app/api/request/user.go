package request

// CreateUser stores request data for POST /users endpoint.
type CreateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
