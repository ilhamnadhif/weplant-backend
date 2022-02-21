package web

// Response

type TransactionProductResponse struct {
	ProductId   string        `json:"product_id"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Description string        `json:"description"`
	Price       int           `json:"price"`
	Quantity    int           `json:"quantity"`
	MainImage   ImageResponse `json:"main_image"`
}
type TransactionActionResponse struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

type TransactionDetailResponse struct {
	Id          string                       `json:"id"`
	CreatedAt   int                          `json:"created_at"`
	UpdatedAt   int                          `json:"updated_at"`
	PaymentType string                       `json:"payment_type"`
	Status      string                       `json:"status"`
	Actions     []TransactionActionResponse  `json:"actions"`
	TotalPrice  int                          `json:"total_price"`
	Products    []TransactionProductResponse `json:"products"`
	Address     AddressResponse              `json:"address"`
}

type TransactionResponse struct {
	CustomerId   string                      `json:"customer_id"`
	Transactions []TransactionDetailResponse `json:"transactions"`
}

// Request

type TransactionCreateRequest struct {
	CreatedAt  int                   `json:"created_at"`
	UpdatedAt  int                   `json:"updated_at"`
	CustomerId string                `json:"customer_id"`
	Address    *AddressCreateRequest `json:"address"`
}

type TransactionCreateRequestResponse struct {
	CreatedAt   int                         `json:"created_at"`
	UpdatedAt   int                         `json:"updated_at"`
	PaymentType string                      `json:"payment_type"`
	Status      string                      `json:"status"`
	Actions     []TransactionActionResponse `json:"actions"`
	TotalPrice  int                         `json:"total_price"`
	Address     AddressCreateRequest        `json:"address"`
}
