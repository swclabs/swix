package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	articleContainer "swclabs/swipex/internal/apis/container/article"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/repos/news"
	articleService "swclabs/swipex/internal/core/service/article"
	"swclabs/swipex/pkg/lib/logger"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func TestPhoneCarousel(t *testing.T) {
	var (
		repos      = news.Mock{}
		service    = articleService.New(nil, nil, &repos)
		controller = articleContainer.NewController(service)
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
		cards = []dtos.CardArticle{
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
		news []entity.News
	)

	for idx, card := range cards {
		json, _ := json.Marshal(card)
		news = append(news, entity.News{
			ID:       int64(idx),
			Category: "mac",
			Header:   "GetByUserID to know your iPhone.",
			Body:     string(json),
			Created:  time.Now().UTC(),
		})
	}

	repos.On("GetMany", context.Background(), "mac", 6).Return(news, nil)

	e.GET("/news", controller.GetNews)
	req := httptest.NewRequest(http.MethodGet, "/news?category=mac&limit=6", nil)
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
