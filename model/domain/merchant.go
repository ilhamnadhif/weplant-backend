package domain

type Merchant struct {
	InitModel
	Email     string   `bson:"email,omitempty"`
	Password  string   `bson:"password,omitempty"`
	Name      string   `bson:"name,omitempty"`
	Phone     string   `bson:"phone,omitempty"`
	MainImage *Image   `bson:"main_image,omitempty"`
	Address   *Address `bson:"address,omitempty"`
}
