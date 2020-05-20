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

// FindUsers represents query params for searching users
type FindUsers struct {
	Limit    int    `form:"limit"`
	BeforeID int    `form:"before_id"`
	AfterID  int    `form:"after_id"`
	Sort     string `form:"sort"`
	Name     string `form:"name"`
	Surname  string `form:"surname"`
	Gender   string `form:"gender"`
	Address  string `form:"address"`
	MinAge   int    `form:"min_age"`
	MaxAge   int    `form:"max_age"`
}
