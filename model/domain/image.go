package domain

type Image struct {
	InitModel
	FileName string `bson:"file_name,omitempty"`
	URL      string `bson:"url,omitempty"`
}
