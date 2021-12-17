package domain

type Admin struct {
	InitModel
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}
