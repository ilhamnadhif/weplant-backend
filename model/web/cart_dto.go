package web

type CartProductResponse struct {
	CreatedAt   int            `json:"created_at"`
	UpdatedAt   int            `json:"updated_at"`
	ProductId   string         `json:"product_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	MainImage   *ImageResponse `json:"main_image"`
	Quantity    int            `json:"quantity"`
}

type CartResponse struct {
	CustomerId string                 `json:"user_id"`
	Total      int                    `json:"total"`
	Products   []*CartProductResponse `json:"products"`
}

type CartProductCreateRequest struct {
	CreatedAt  int
	UpdatedAt  int
	CustomerId string
	ProductId  string `json:"product_id"`
	Quantity   int
}

type CartProductUpdateRequest struct {
	UpdatedAt  int
	CustomerId string
	ProductId  string
	Quantity   int `json:"quantity"`
}
