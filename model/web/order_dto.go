package web

type OrderProductResponse struct {
	Id          string           `json:"id"`
	CreatedAt   int              `json:"created_at"`
	UpdatedAt   int              `json:"updated_at"`
	ProductId   string           `json:"product_id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       int              `json:"price"`
	MainImage   *ImageResponse   `json:"main_image"`
	Quantity    int              `json:"quantity"`
	HasDone     *bool            `json:"has_done"`
	Address     *AddressResponse `json:"address"`
}

type OrderResponse struct {
	CustomerId string                  `json:"user_id"`
	Products   []*OrderProductResponse `json:"products"`
}

type OrderProductCreateRequest struct {
	CreatedAt  int
	UpdatedAt  int
	CustomerId string
	Address    *AddressCreateRequest `json:"address"`
}
