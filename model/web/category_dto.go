package web

// Response

type CategoryDetailResponse struct {
	Id        string                  `json:"id"`
	CreatedAt int                     `json:"created_at"`
	UpdatedAt int                     `json:"updated_at"`
	Name      string                  `json:"name"`
	Slug      string                  `json:"slug"`
	Products  []ProductSimpleResponse `json:"products"`
}

type CategorySimpleResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Request

type CategoryCreateRequest struct {
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
}

type CategoryCreateRequestResponse struct {
	Id        string `json:"id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
}
