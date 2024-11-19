package article

import (
	"net/http"
	"strconv"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/service/article"
	"github.com/swclabs/swipex/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

var _ = app.Controller(NewController)

// NewController creates a new Article object
func NewController(service article.IArticle) IController {
	return &Controller{
		Services: service,
	}
}

// IController interface for article controller
type IController interface {
	UploadNews(c echo.Context) error
	UpdateNewsImage(c echo.Context) error
	GetNews(c echo.Context) error

	GetComment(c echo.Context) error
	UploadComment(c echo.Context) error
}

// Controller struct implementation of IArticle
type Controller struct {
	Services article.IArticle
}

// GetNews .
// @Description get news
// @Tags news
// @Accept json
// @Produce json
// @Param category query string true "category of news"
// @Param limit query number true "limit of cards carousel"
// @Success 200 {object} dtos.News
// @Router /news [GET]
func (p *Controller) GetNews(c echo.Context) error {
	category := c.QueryParam("category")
	limit := c.QueryParam("limit")
	if category == "" || limit == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'category' or 'limit' field",
		})
	}

	_limit, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}

	carousel, err := p.Services.GetNews(c.Request().Context(), category, _limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, *carousel)
}

// UpdateNewsImage .
// @Description update news image
// @Tags news
// @Accept json
// @Produce json
// @Param img formData file true "image of news"
// @Param id formData string true "news identifier"
// @Success 200 {object} dtos.OK
// @Router /news/image [PUT]
func (p *Controller) UpdateNewsImage(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	id, err := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "'id' field invalid",
		})
	}

	if err := p.Services.UploadNewsImage(c.Request().Context(), id, file); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "upload image of collection successfully",
	})
}

// UploadNews .
// @Description create news
// @Tags news
// @Accept json
// @Produce json
// @Param collection body dtos.NewsDTO true "news Request"
// @Success 201 {object} dtos.OK
// @Router /news [POST]
func (p *Controller) UploadNews(c echo.Context) error {
	var news dtos.NewsDTO
	if err := c.Bind(&news); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&news); _valid != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: _valid.Error(),
		})
	}
	_, err := p.Services.UploadNews(c.Request().Context(), news)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "collection uploaded successfully",
	})
}

// GetComment .
// @Description get all comments of product
// @Tags comments
// @Accept json
// @Produce json
// @Param product_id query string true "id of products"
// @Success 200 {object} []model.Comment
// @Router /comment [GET]
func (p *Controller) GetComment(c echo.Context) error {
	product_id, err := strconv.Atoi(c.QueryParam("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Missing 'product_id' required",
		})
	}

	comments, err := p.Services.GetComment(c.Request().Context(), int64(product_id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, comments)
}

// UploadComment .
// @Description create comment into products
// @Tags comment
// @Accept json
// @Produce json
// @Param banner body dtos.Comment true "comment data request"
// @Success 201 {object} dtos.OK
// @Router /comment [POST]
func (p *Controller) UploadComment(c echo.Context) error {
	var cmt dtos.Comment
	if err := c.Bind(&cmt); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&cmt); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := p.Services.UploadComment(c.Request().Context(), cmt); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "your comment has been uploaded successfully",
	})
}
