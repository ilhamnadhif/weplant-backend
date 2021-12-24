package web

type CategoryResponse struct {
	Id        string         `json:"id"`
	CreatedAt int            `json:"created_at"`
	UpdatedAt int            `json:"updated_at"`
	Name      string         `json:"name"`
	MainImage *ImageResponse `json:"main_image"`
}

type CategoryResponseWithProduct struct {
	Id        string                `json:"id"`
	CreatedAt int                   `json:"created_at"`
	UpdatedAt int                   `json:"updated_at"`
	Name      string                `json:"name"`
	MainImage *ImageResponse        `json:"main_image"`
	Products  []*ProductResponseAll `json:"products"`
}

type CategoryCreateRequest struct {
	CreatedAt int                 `json:"created_at"`
	UpdatedAt int                 `json:"updated_at"`
	Name      string              `json:"name"`
	MainImage *ImageCreateRequest `json:"main_image"`
}

type CategoryUpdateRequest struct {
	Id        string `json:"id"`
	UpdatedAt int    `json:"updated_at"`
	Name      string `json:"name"`
}

type CategoryUpdateImageRequest struct {
	Id        string              `json:"id"`
	UpdatedAt int                 `json:"updated_at"`
	MainImage *ImageUpdateRequest `json:"main_image"`
}
