package web

type MerchantResponse struct {
	Id        string           `json:"id"`
	CreatedAt int              `json:"created_at"`
	UpdatedAt int              `json:"updated_at"`
	Email     string           `json:"email"`
	Name      string           `json:"name"`
	Phone     string           `json:"phone"`
	MainImage *ImageResponse   `json:"main_image"`
	Address   *AddressResponse `json:"address"`
}

type MerchantCreateRequest struct {
	CreatedAt int                   `json:"created_at"`
	UpdatedAt int                   `json:"updated_at"`
	Email     string                `json:"email,omitempty"`
	Password  string                `json:"password,omitempty"`
	Name      string                `json:"name,omitempty"`
	Phone     string                `json:"phone,omitempty"`
	MainImage *ImageCreateRequest   `json:"main_image,omitempty"`
	Address   *AddressCreateRequest `json:"address,omitempty"`
}

type MerchantUpdateRequest struct {
	Id        string                `json:"id"`
	UpdatedAt int                   `json:"updated_at"`
	Name      string                `json:"name,omitempty"`
	Phone     string                `json:"phone,omitempty"`
	Address   *AddressUpdateRequest `json:"address,omitempty"`
}

type MerchantUpdateImageRequest struct {
	Id        string              `json:"id"`
	UpdatedAt int                 `json:"updated_at"`
	MainImage *ImageUpdateRequest `json:"main_image,omitempty"`
}
