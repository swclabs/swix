// Authors: Duc Hung Ho @kieranhoo
// Description: Plugin for tasks package, run worker tasks in memory

package resolver

import (
	"context"
	"log"
	"mime/multipart"
	"sync"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/repo"
	"swclabs/swipe-api/pkg/cloud"
	"swclabs/swipe-api/pkg/tools"
)

var ImagePool *tools.Pool

func StarUserImageHandler(concurrent int) {
	ImagePool = tools.NewPool()
	img := UserImage{}
	tools.NewTask(ImagePool, img.UserImageHandler, concurrent).Run()
}

type UserImage struct {
	Email string
	File  multipart.File
}

func (image *UserImage) UserImageHandler(data <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _data := range data {
		if imageInfo, ok := _data.(UserImage); ok {
			resp, err := cloud.UpdateImages(cloud.Connection(), imageInfo.File)
			if err != nil {
				log.Fatal(err)
			}
			if err := repo.NewUsers().SaveInfo(
				context.TODO(),
				&domain.User{
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
