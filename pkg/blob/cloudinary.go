package blob

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"sync"

	"swclabs/swipecore/internal/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary
var lockCld = &sync.Mutex{}

func Connection() *cloudinary.Cloudinary {
	if cld == nil {
		lockCld.Lock()
		defer lockCld.Unlock()
		if cld == nil {
			cldLocal, err := cloudinary.NewFromURL(config.CloudinaryUrl)
			if err != nil {
				log.Fatal(err)
			}
			cld = cldLocal
			fmt.Println("Cloudinary connected !!")
			return cld
		} else {
			return cld
		}
	}
	return cld
}

func UploadImages(cld *cloudinary.Cloudinary, file interface{}) (*uploader.UploadResult, error) {
	var ctx = context.Background()
	return UploadImagesWithContext(ctx, cld, file)
}

func UploadImagesWithContext(ctx context.Context, cld *cloudinary.Cloudinary, file interface{}) (*uploader.UploadResult, error) {
	updateResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "auto",
		Folder:       "swc-storage",
	})
	return updateResult, err
}

func UploadFile(ctx context.Context, cld *cloudinary.Cloudinary, fileHeader *multipart.FileHeader) (url string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	resp, err := UploadImagesWithContext(ctx, cld, file)
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
