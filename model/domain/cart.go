package domain

type CartProduct struct {
	CreatedAt int    `bson:"created_at,omitempty"`
	UpdatedAt int    `bson:"updated_at,omitempty"`
	ProductId string `bson:"product_id,omitempty"`
	Quantity  int    `bson:"quantity,omitempty"`
}

type Cart struct {
	SubTotal int            `bson:"sub_total"`
	Products []*CartProduct `bson:"products,omitempty"`
}
