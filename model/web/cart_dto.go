package web

type CartProductResponse struct {
	ProductId   string         `json:"product_id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Quantity    int            `json:"quantity"`
	SubTotal    int            `json:"sub_total"`
	MainImage   *ImageResponse `json:"main_image"`
}

type CartResponse struct {
	CustomerId string                 `json:"user_id"`
	TotalPrice int                    `json:"total_price"`
	Products   []*CartProductResponse `json:"products"`
}

type CartProductCreateRequest struct {
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

type CartProductUpdateRequest struct {
	UpdatedAt  int    `json:"updated_at"`
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}
