package domain

type ProductCategory struct {
	CategoryId string `bson:"category_id,omitempty"`
}

type Product struct {
	InitModel
	MerchantId  string             `bson:"merchant_id,omitempty"`
	IsActive    *bool              `bson:"is_active,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Price       int                `bson:"price,omitempty"`
	Quantity    int                `bson:"quantity,omitempty"`
	MainImage   *Image             `bson:"main_image,omitempty"`
	Images      []*Image           `bson:"images,omitempty"`
	Categories  []*ProductCategory `bson:"categories,omitempty"`
}
