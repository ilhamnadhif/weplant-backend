package web

type CustomerResponse struct {
	Id        string `json:"id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
}

type CustomerCreateRequest struct {
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type CustomerUpdateRequest struct {
	Id        string `json:"id"`
	UpdatedAt int    `json:"updated_at"`
	UserName  string `json:"user_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
