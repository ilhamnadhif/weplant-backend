package web

type ManageOrderProductResponse struct {
	Id          string         `json:"id"`
	CreatedAt   int            `json:"created_at"`
	UpdatedAt   int            `json:"updated_at"`
	ProductId   string         `json:"product_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	MainImage   *ImageResponse `json:"main_image"`
	Quantity    int            `json:"quantity"`
}

type ManageOrderResponse struct {
	MerchantId string                        `json:"user_id"`
	Products   []*ManageOrderProductResponse `json:"products"`
}
