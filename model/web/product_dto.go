package web

// Response

type ProductDetailResponse struct {
	Id          string                   `json:"id"`
	CreatedAt   int                      `json:"created_at"`
	UpdatedAt   int                      `json:"updated_at"`
	Name        string                   `json:"name"`
	Slug        string                   `json:"slug"`
	Description string                   `json:"description"`
	Price       int                      `json:"price"`
	Stock       int                      `json:"stock"`
	MainImage   ImageResponse            `json:"main_image"`
	Images      []ImageResponse          `json:"images"`
	Categories  []CategorySimpleResponse `json:"categories"`
	Merchant    MerchantSimpleResponse   `json:"merchant"`
}

type ProductSimpleResponse struct {
	Id          string        `json:"id"`
	MerchantId  string        `json:"merchant_id"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Description string        `json:"description"`
	Price       int           `json:"price"`
	Stock       int           `json:"stock"`
	MainImage   ImageResponse `json:"main_image"`
}

type MetadataPaginationResponse struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalData   int `json:"total_data"`
}

type ProductFindAllResponse struct {
	Products []ProductSimpleResponse    `json:"products"`
	Metadata MetadataPaginationResponse `json:"metadata"`
}

// Request

type ProductCategoryCreateRequest struct {
	CategoryId string `json:"category_id"`
}

type ProductCreateRequest struct {
	CreatedAt   int                            `json:"created_at"`
	UpdatedAt   int                            `json:"updated_at"`
	MerchantId  string                         `json:"merchant_id"`
	Name        string                         `json:"name"`
	Slug        string                         `json:"slug"`
	Description string                         `json:"description"`
	Price       int                            `json:"price"`
	Stock       int                            `json:"stock"`
	MainImage   *ImageCreateRequest            `json:"main_image"`
	Images      []ImageCreateRequest           `json:"images"`
	Categories  []ProductCategoryCreateRequest `json:"categories"`
}

type ProductCreateRequestResponse struct {
	Id          string                         `json:"id"`
	CreatedAt   int                            `json:"created_at"`
	UpdatedAt   int                            `json:"updated_at"`
	MerchantId  string                         `json:"merchant_id"`
	Name        string                         `json:"name"`
	Slug        string                         `json:"slug"`
	Description string                         `json:"description"`
	Price       int                            `json:"price"`
	Stock       int                            `json:"stock"`
	MainImage   ImageResponse                  `json:"main_image"`
	Images      []ImageResponse                `json:"images"`
	Categories  []ProductCategoryCreateRequest `json:"categories"`
}

type ProductCategoryUpdateRequest struct {
	CategoryId string `json:"category_id"`
}

type ProductUpdateRequest struct {
	Id          string                         `json:"id"`
	UpdatedAt   int                            `json:"updated_at"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	Price       int                            `json:"price"`
	Stock       int                            `json:"stock"`
	Categories  []ProductCategoryUpdateRequest `json:"categories"`
}

type ProductUpdateImageRequest struct {
	Id        string              `json:"id"`
	UpdatedAt int                 `json:"updated_at"`
	MainImage *ImageUpdateRequest `json:"main_image"`
}

type ProductUpdateImageRequestResponse struct {
	Id        string        `json:"id"`
	UpdatedAt int           `json:"updated_at"`
	MainImage ImageResponse `json:"main_image"`
}
