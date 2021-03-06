package web

// Response

type CustomerResponse struct {
	Id        string        `json:"id"`
	CreatedAt int           `json:"created_at"`
	UpdatedAt int           `json:"updated_at"`
	Email     string        `json:"email"`
	UserName  string        `json:"user_name"`
	Phone     string        `json:"phone"`
	MainImage ImageResponse `json:"main_image"`
}

// Request

type CustomerCreateRequest struct {
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
}

type CustomerUpdateRequest struct {
	Id        string `json:"id"`
	UpdatedAt int    `json:"updated_at"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
}

type CustomerUpdateImageRequest struct {
	Id        string              `json:"id"`
	UpdatedAt int                 `json:"updated_at"`
	MainImage *ImageUpdateRequest `json:"main_image"`
}

type CustomerUpdateImageRequestResponse struct {
	Id        string        `json:"id"`
	UpdatedAt int           `json:"updated_at"`
	MainImage ImageResponse `json:"main_image"`
}
