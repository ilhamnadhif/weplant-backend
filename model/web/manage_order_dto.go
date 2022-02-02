package web

type ManageOrderProductResponse struct {
	Id          string           `json:"id"`
	CreatedAt   int              `json:"created_at"`
	UpdatedAt   int              `json:"updated_at"`
	ProductId   string           `json:"product_id"`
	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	Price       int              `json:"price"`
	Quantity    int              `json:"quantity"`
	SubTotal    int              `json:"sub_total"`
	MainImage   *ImageResponse   `json:"main_image"`
	Address     *AddressResponse `json:"address"`
}

type ManageOrderResponse struct {
	MerchantId string                        `json:"user_id"`
	Products   []*ManageOrderProductResponse `json:"products"`
}
