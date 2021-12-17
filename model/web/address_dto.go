package web

type AddressResponse struct {
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	Country    string  `json:"country"`
	PostalCode int     `json:"postal_code"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type AddressCreateRequest struct {
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	Country    string  `json:"country"`
	PostalCode int     `json:"postal_code"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type AddressUpdateRequest struct {
	Id         string
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Province   string  `json:"province"`
	Country    string  `json:"country"`
	PostalCode int     `json:"postal_code"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
