// Package blob connect to blob storage service
package blob

import (
	"context"
	"fmt"
	"mime/multipart"
	"swclabs/swipex/internal/config"
	"sync"

	"swclabs/swipex/pkg/lib/logger"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.uber.org/fx"
)

var cld *cloudinary.Cloudinary
var lockCld = &sync.Mutex{}

// Connection creates a new cloudinary connection.
func Connection() IBlobStorage {
	return &Storage{
		Conn: cld,
	}
}

// Storage struct implements IBlobStorage interface
type Storage struct {
	Conn *cloudinary.Cloudinary
}

// New creates a new cloudinary connection.
func New(lc fx.Lifecycle) IBlobStorage {
	var err error
	if cld == nil {
		lockCld.Lock()
		defer lockCld.Unlock()
		if cld == nil {
			cld, err = cloudinary.NewFromURL(config.CloudinaryURL)
		}
	}
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			if err != nil {
				return err
			}
			logger.Info(fmt.Sprintf(
				"%s to %s", logger.Green.Add("Connect"), logger.Blue.Add("Cloudinary")),
			)
			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Info(fmt.Sprintf(
				"%s %s connection", logger.Green.Add("Closed"), logger.Blue.Add("Cloudinary")),
			)
			return nil
		},
	})
	return &Storage{
		Conn: cld,
	}
}

// UploadImages upload images to cloudinary
func (blob *Storage) UploadImages(file interface{}) (UploadResult, error) {
	var ctx = context.Background()
	return blob.UploadImagesWithContext(ctx, file)
}

// UploadImagesWithContext upload images to cloudinary with context
func (blob *Storage) UploadImagesWithContext(ctx context.Context, file interface{}) (UploadResult, error) {
	updateResult, err := blob.Conn.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "auto",
		Folder:       "swc-storage",
	})
	return updateResult, err
}

// UploadFile upload file to cloudinary
func (blob *Storage) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (url string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	resp, err := blob.UploadImagesWithContext(ctx, file)
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
