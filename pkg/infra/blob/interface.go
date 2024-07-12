package blob

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type (
	// UploadResult struct for upload result
	UploadResult *uploader.UploadResult
)

// IBlobStorage interface for blob storage
type IBlobStorage interface {
	UploadImages(file interface{}) (UploadResult, error)
	UploadImagesWithContext(ctx context.Context, file interface{}) (UploadResult, error)
	UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (url string, err error)
}
