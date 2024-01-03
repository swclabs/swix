// Authors: Duc Hung Ho @ikierans
// Description: Plugin for tasks package, run worker tasks in memory

package plugin

import (
	"log"
	"mime/multipart"
	"swclabs/swiftcart/internal/domain"
	"swclabs/swiftcart/internal/repo"
	"swclabs/swiftcart/pkg/cloud"
	"sync"
)

var ImagePool *Pool

func StartImageHandler(concurrent int) {
	ImagePool = NewPool()
	img := Image{}
	NewTask(ImagePool, img.ImageHandler, concurrent).Run()
}

type Image struct {
	Email string
	File  multipart.File
}

func (image *Image) ImageHandler(data <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _data := range data {
		if imageInfo, ok := _data.(Image); ok {
			resp, err := cloud.UpdateImages(cloud.Connection(), imageInfo.File)
			if err != nil {
				log.Fatal(err)
			}
			if err := repo.NewUsers().SaveInfo(&domain.User{
				Email: imageInfo.Email,
				Image: resp.SecureURL,
			}); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("can't update images: unknown image type")
		}
	}
}
