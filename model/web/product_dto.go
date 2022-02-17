package web

type ProductResponse struct {
	Id          string            `json:"id"`
	CreatedAt   int               `json:"created_at"`
	UpdatedAt   int               `json:"updated_at"`
	MerchantId  string            `json:"merchant_id"`
	Name        string            `json:"name"`
	Slug        string            `json:"slug"`
	Description string            `json:"description"`
	Price       int               `json:"price"`
	Stock       int               `json:"stock"`
	MainImage   *ImageResponse    `json:"main_image"`
	Images      []ImageResponse   `json:"images"`
	Category    *CategoryResponse `json:"category"`
}

type ProductResponseAll struct {
	Id          string         `json:"id"`
	CreatedAt   int            `json:"created_at"`
	UpdatedAt   int            `json:"updated_at"`
	MerchantId  string         `json:"merchant_id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Stock       int            `json:"stock"`
	MainImage   *ImageResponse `json:"main_image"`
}

type ProductCreateRequest struct {
	CreatedAt   int                  `json:"created_at"`
	UpdatedAt   int                  `json:"updated_at"`
	MerchantId  string               `json:"merchant_id"`
	Name        string               `json:"name"`
	Slug        string               `json:"slug"`
	Description string               `json:"description"`
	Price       int                  `json:"price"`
	Stock       int                  `json:"stock"`
	MainImage   *ImageCreateRequest  `json:"main_image"`
	Images      []ImageCreateRequest `json:"images"`
	CategoryId  string               `json:"category_id"`
}

type ProductUpdateRequest struct {
	Id          string `json:"id"`
	UpdatedAt   int    `json:"updated_at"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryId  string `json:"category_id"`
}

type ProductUpdateImageRequest struct {
	Id        string              `json:"id"`
	UpdatedAt int                 `json:"updated_at"`
	MainImage *ImageUpdateRequest `json:"main_image"`
}
