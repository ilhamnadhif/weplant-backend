package web

type ImageResponse struct {
	Id        string `json:"id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	FileName  string `json:"file_name"`
	URL       string `json:"url"`
}

type ImageCreateRequest struct {
	FileName string `json:"file_name"`
}

type ImageUpdateRequest struct {
	FileName string `json:"file_name"`
}
