package web

type TransactionProductResponse struct {
	ProductId   string         `json:"product_id"`
	Name        string         `json:"name"`
	Slug        string         `json:"slug"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Quantity    int            `json:"quantity"`
	SubTotal    int            `json:"sub_total"`
	MainImage   *ImageResponse `json:"main_image"`
}

type TransactionDetailResponse struct {
	Id         string `json:"id"`
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
	Status     string `json:"status"`
	QRCode     string `json:"qr_code"`
	TotalPrice int    `json:"total_price"`
	Products   []*TransactionProductResponse
	Address    *AddressResponse `json:"address"`
}

type TransactionResponse struct {
	CustomerId   string                       `json:"customer_id"`
	Transactions []*TransactionDetailResponse `json:"transactions"`
}

type TransactionCreateRequest struct {
	CreatedAt  int                   `json:"created_at"`
	UpdatedAt  int                   `json:"updated_at"`
	CustomerId string                `json:"customer_id"`
	Address    *AddressCreateRequest `json:"address"`
}

type TransactionCreateRequestResponse struct {
	CreatedAt  int                   `json:"created_at"`
	UpdatedAt  int                   `json:"updated_at"`
	CustomerId string                `json:"customer_id"`
	Status     string                `json:"status"`
	QRCode     string                `json:"qr_code"`
	Address    *AddressCreateRequest `json:"address"`
}