package request

// CreateUser stores request data for POST /v1/users endpoint.
type CreateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

// UpdateUser stores request data for PU /v1/users/:user_id endpoint.
type UpdateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
