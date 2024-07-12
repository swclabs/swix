// Package blob connect to blob storage service
package blob

import (
	"context"
	"fmt"
	"mime/multipart"
	"sync"

	"swclabs/swipecore/internal/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.uber.org/fx"
)

var cld *cloudinary.Cloudinary
var lockCld = &sync.Mutex{}

func Connection() IBlobStorage {
	return &BlobStorage{
		Conn: cld,
	}
}

type BlobStorage struct {
	Conn *cloudinary.Cloudinary
}

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
			fmt.Printf("[SWIPE]-v%s ===============> connect to cloudinary\n", config.Version)
			return nil
		},
		OnStop: func(_ context.Context) error {
			fmt.Printf("[SWIPE]-v%s ===============> closed cloudinary connection\n", config.Version)
			return nil
		},
	})
	return &BlobStorage{
		Conn: cld,
	}
}

func (blob *BlobStorage) UploadImages(file interface{}) (UploadResult, error) {
	var ctx = context.Background()
	return blob.UploadImagesWithContext(ctx, file)
}

func (blob *BlobStorage) UploadImagesWithContext(ctx context.Context, file interface{}) (UploadResult, error) {
	updateResult, err := blob.Conn.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "auto",
		Folder:       "swc-storage",
	})
	return updateResult, err
}

func (blob *BlobStorage) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (url string, err error) {
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
