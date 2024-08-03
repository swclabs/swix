package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/repository/collections"
	"swclabs/swipecore/internal/core/service/article"
	"swclabs/swipecore/internal/webapi/controller"
	"swclabs/swipecore/pkg/lib/logger"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func TestPhoneCarousel(t *testing.T) {
	var (
		repos      = collections.Mock{}
		service    = article.New(nil, &repos)
		controller = controller.NewArticle(service)
		e          = echo.New()

		content = []dtos.CardContent{
			{
				Content: "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought.",
				Src:     "/img/posts/8.jpg",
			},
			{
				Content: "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought.",
				Src:     "/img/posts/8.jpg",
			},
			{
				Content: "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought.",
				Src:     "/img/posts/8.jpg",
			},
		}
		cards = []dtos.CardInArticle{
			{
				Category: "Artificial Intelligence",
				Title:    "You can do more with AI.",
				Src:      "/img/posts/1.jpg",
				Content:  content,
			},
			{
				Category: "Productivity",
				Title:    "Enhance your productivity.",
				Src:      "/img/posts/2.jpg",
				Content:  content,
			},
			{
				Category: "Product",
				Title:    "Launching the new Apple Vision Pro.",
				Src:      "/img/posts/3.jpg",
				Content:  content,
			},
			{
				Category: "Product",
				Title:    "Maps for your iPhone 15 Pro Max.",
				Src:      "/img/posts/4.jpg",
				Content:  content,
			},
			{
				Category: "iOS",
				Title:    "Photography just got better.",
				Src:      "/img/posts/5.jpg",
				Content:  content,
			},
			{
				Category: "Hiring",
				Title:    "Hiring for a Staff Software Engineer.",
				Src:      "/img/posts/6.jpg",
				Content:  content,
			},
		}
		collection []entity.Collection
	)

	for idx, card := range cards {
		json, _ := json.Marshal(card)
		collection = append(collection, entity.Collection{
			ID:       int64(idx),
			Position: "mac_1",
			Headline: "Get to know your iPhone.",
			Body:     string(json),
			Created:  time.Now().UTC(),
		})
	}

	repos.On("GetMany", context.Background(), "mac_1", 6).Return(collection, nil)

	e.GET("/collections", controller.GetArticleData)
	req := httptest.NewRequest(http.MethodGet, "/collections?position=mac_1&limit=6", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	responseBody := rr.Body.Bytes()

	var body dtos.Article
	if err := json.Unmarshal(responseBody, &body); err != nil {
		t.Error(err)
	}
	file, err := os.Create("./phone_out.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	logger := logger.Write(file)
	logger.Info("Response body", zap.Any("body", body), zap.Int("status", rr.Code))
}
