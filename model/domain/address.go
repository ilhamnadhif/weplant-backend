package domain

type Address struct {
	Address    string  `bson:"address,omitempty"`
	City       string  `bson:"city,omitempty"`
	Province   string  `bson:"province,omitempty"`
	Country    string  `bson:"country,omitempty"`
	PostalCode string  `bson:"postal_code,omitempty"`
	Latitude   float64 `bson:"latitude,omitempty"`
	Longitude  float64 `bson:"longitude,omitempty"`
}
