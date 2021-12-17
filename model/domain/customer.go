package domain

type Customer struct {
	InitModel
	Email    string   `bson:"email,omitempty"`
	Password string   `bson:"password,omitempty"`
	UserName string   `bson:"user_name,omitempty"`
	Phone    string   `bson:"phone,omitempty"`
	Address  *Address `bson:"address,omitempty"`
}
