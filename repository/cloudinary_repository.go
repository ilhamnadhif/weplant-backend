package repository

import (
	"context"
)

type CloudinaryRepository interface {
	UploadImage(ctx context.Context, filename string, image interface{}) (string, error)
	DeleteImage(ctx context.Context, filename string) error
}
