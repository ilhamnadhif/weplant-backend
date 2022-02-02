package web

type ImageResponse struct {
	Id       string `json:"id"`
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

type ImageCreateRequest struct {
	FileName string      `json:"file_name"`
	URL      interface{} `json:"url"`
}

type ImageUpdateRequest struct {
	FileName string      `json:"file_name"`
	URL      interface{} `json:"url"`
}
