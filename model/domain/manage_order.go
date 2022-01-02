package domain

type ManageOrderProduct struct {
	CreatedAt int      `bson:"created_at,omitempty"`
	UpdatedAt int      `bson:"updated_at,omitempty"`
	ProductId string   `bson:"product_id,omitempty"`
	Price     int      `bson:"price,omitempty"`
	Quantity  int      `bson:"quantity,omitempty"`
	Address   *Address `bson:"address,omitempty"`
}
