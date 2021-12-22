package web

type ProductResponse struct {
	Id          string              `json:"id"`
	CreatedAt   int                 `json:"created_at"`
	UpdatedAt   int                 `json:"updated_at"`
	MerchantId  string              `json:"merchant_id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       int                 `json:"price"`
	Quantity    int                 `json:"quantity"`
	MainImage   *ImageResponse      `json:"main_image"`
	Images      []*ImageResponse    `json:"images"`
	Categories  []*CategoryResponse `json:"categories"`
}

type ProductResponseAll struct {
	Id          string         `json:"id"`
	CreatedAt   int            `json:"created_at"`
	UpdatedAt   int            `json:"updated_at"`
	MerchantId  string         `json:"merchant_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Quantity    int            `json:"quantity"`
	MainImage   *ImageResponse `json:"main_image"`
}

type ProductCategoryCreateRequest struct {
	CategoryId string `json:"category_id"`
}

type ProductCreateRequest struct {
	MerchantId  string                          `json:"merchant_id"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Price       int                             `json:"price"`
	Quantity    int                             `json:"quantity"`
	MainImage   *ImageCreateRequest             `json:"main_image"`
	Categories  []*ProductCategoryCreateRequest `json:"categories"`
}

type ProductCategoryUpdateRequest struct {
	CategoryId string `json:"category_id"`
}

type ProductUpdateRequest struct {
	Id          string                          `json:"id"`
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	Price       int                             `json:"price"`
	Quantity    int                             `json:"quantity"`
	Categories  []*ProductCategoryUpdateRequest `json:"categories"`
}

//type ProductSetIsActiveRequest struct {
//	Id       string
//	IsActive bool `json:"is_active"`
//}
