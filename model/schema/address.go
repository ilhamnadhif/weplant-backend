package schema

type Address struct {
	Address    string `bson:"address,omitempty"`
	City       string `bson:"city,omitempty"`
	Province   string `bson:"province,omitempty"`
	PostalCode string `bson:"postal_code,omitempty"`
}
