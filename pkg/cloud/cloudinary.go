package cloud

import (
	"context"
	"fmt"
	"log"
	"sync"

	"swclabs/swipe-api/internal/config"

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
	updateResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "auto",
		Folder:       "swc-storage",
	})
	return updateResult, err
}
