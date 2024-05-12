package controller

import (
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/service"
	"swclabs/swipecore/pkg/tools/valid"
	"swclabs/swipecore/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type IPosts interface {
	GetNewsletter(c echo.Context) error
	UploadNewsletter(c echo.Context) error
}

type Posts struct {
	Services domain.IPostsService
}

func NewPosts() IPosts {
	return &Posts{
		Services: service.NewPost(),
	}
}

// GetNewsletter
// @Description Get Product Newsletter
// @Tags posts
// @Accept json
// @Produce json
// @Param limit query int true "limit number of newsletter"
// @Success 200 {object} domain.NewsletterListResponse
// @Router /newsletters [GET]
func (p *Posts) GetNewsletter(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	newsletter, err := p.Services.GetNewsletter(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.NewsletterListResponse{
		Data: newsletter,
	})
}

// UploadNewsletter
// @Description Create newsletter
// @Tags posts
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param img formData file true "image of newsletter"
// @Param product formData domain.Newsletter true "Newsletter Request"
// @Success 200 {object} domain.OK
// @Router /newsletters [POST]
func (p *Posts) UploadNewsletter(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	formData, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// bind json to structure
	var newsletter domain.Newsletter

	if err := mapstructure.Decode(utils.NxN2Nx1(formData.Value), &newsletter); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// check validate struct
	if valid := valid.Validate(&newsletter); valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: valid,
		})
	}
	// call services
	if err := p.Services.UploadNewsletter(c.Request().Context(), newsletter, file); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "upload newsletter successfully",
	})
}
