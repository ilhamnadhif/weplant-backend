package web

// Response

type OrderProductResponse struct {
	Id          string          `json:"id"`
	CreatedAt   int             `json:"created_at"`
	UpdatedAt   int             `json:"updated_at"`
	ProductId   string          `json:"product_id"`
	Name        string          `json:"name"`
	Slug        string          `json:"slug"`
	Description string          `json:"description"`
	Price       int             `json:"price"`
	Quantity    int             `json:"quantity"`
	MainImage   ImageResponse   `json:"main_image"`
	Address     AddressResponse `json:"address"`
}

type OrderResponse struct {
	CustomerId string                 `json:"customer_id"`
	Products   []OrderProductResponse `json:"products"`
}
