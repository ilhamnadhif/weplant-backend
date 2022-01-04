package web

type OrderProductResponse struct {
	ProductId   string         `json:"product_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	MainImage   *ImageResponse `json:"main_image"`
	Quantity    int            `json:"quantity"`
}

type OrderResponse struct {
	Id        string                  `json:"id"`
	CreatedAt int                     `json:"created_at"`
	UpdatedAt int                     `json:"updated_at"`
	Products  []*OrderProductResponse `json:"products"`
	Address   *AddressResponse        `json:"address"`
}

type CustomerOrdersResponse struct {
	CustomerId string           `json:"customer_id"`
	Orders     []*OrderResponse `json:"orders"`
}

type OrderCreateRequest struct {
	CustomerId string                `json:"customer_id"`
	Address    *AddressCreateRequest `json:"address"`
}

//type OrderCreateRequestMidtrans struct {
//	CustomerId string  `json:"customer_id"`
//	Address    string  `json:"address"`
//	City       string  `json:"city"`
//	Province   string  `json:"province"`
//	Country    string  `json:"country"`
//	PostalCode string  `json:"postal_code"`
//	Latitude   float64 `json:"latitude"`
//	Longitude  float64 `json:"longitude"`
//}
