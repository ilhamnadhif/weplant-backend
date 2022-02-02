package domain

type CartProduct struct {
	ProductId string `bson:"product_id,omitempty"`
	Quantity  int    `bson:"quantity,omitempty"`
}
