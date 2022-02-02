package web

type MerchantResponse struct {
	Id        string                 `json:"id"`
	CreatedAt int                    `json:"created_at"`
	UpdatedAt int                    `json:"updated_at"`
	Email     string                 `json:"email"`
	Name      string                 `json:"name"`
	Slug      string                 `json:"slug"`
	Phone     string                 `json:"phone"`
	Balance   int64                  `json:"balance"`
	MainImage *ImageResponse         `json:"main_image"`
	Address   *AddressResponse       `json:"address"`
	Products  []*ProductResponseAll `json:"products"`
}

type MerchantCreateRequest struct {
	CreatedAt int                   `json:"created_at"`
	UpdatedAt int                   `json:"updated_at"`
	Email     string                `json:"email"`
	Password  string                `json:"password"`
	Name      string                `json:"name"`
	Slug      string                `json:"slug"`
	Phone     string                `json:"phone"`
	MainImage *ImageCreateRequest   `json:"main_image"`
	Address   *AddressCreateRequest `json:"address"`
}

type MerchantUpdateRequest struct {
	Id        string                `json:"id"`
	UpdatedAt int                   `json:"updated_at"`
	Name      string                `json:"name"`
	Phone     string                `json:"phone"`
	Address   *AddressUpdateRequest `json:"address"`
}

type MerchantUpdateImageRequest struct {
	Id        string              `json:"id"`
	UpdatedAt int                 `json:"updated_at"`
	MainImage *ImageUpdateRequest `json:"main_image"`
}
