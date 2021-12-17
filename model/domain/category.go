package domain

type Category struct {
	InitModel
	Name  string `bson:"name,omitempty"`
	Image *Image `bson:"image,omitempty"`
}
