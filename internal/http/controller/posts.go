package controller

import (
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/service"
	"swclabs/swipecore/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

type IPosts interface {
	UploadCollections(c echo.Context) error
	UpdateCollectionsImage(c echo.Context) error
	GetSlicesOfCollections(c echo.Context) error

	UploadHeadlineBanner(c echo.Context) error
	GetSliceOfHeadlineBanner(c echo.Context) error
}

func NewPosts() IPosts {
	return &Posts{
		Services: service.NewPost(),
	}
}

type Posts struct {
	Services domain.IPostsService
}

// GetSliceOfHeadlineBanner implements IPosts.
func (p *Posts) GetSliceOfHeadlineBanner(c echo.Context) error {
	panic("unimplemented")
}

// UploadHeadlineBanner implements IPosts.
func (p *Posts) UploadHeadlineBanner(c echo.Context) error {
	panic("unimplemented")
}

// UploadCollections
// @Description Create collections
// @Tags posts
// @Accept json
// @Produce json
// @Param collection body domain.CollectionTypeSwagger true "collections Request"
// @Success 201 {object} domain.CollectionUploadRes
// @Router /collections [POST]
func (p *Posts) UploadCollections(c echo.Context) error {
	var cardBanner domain.CollectionType
	if err := c.Bind(&cardBanner); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(cardBanner); _valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
	}
	id, err := p.Services.UploadCollections(c.Request().Context(), cardBanner)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.CollectionUploadRes{
		Msg: "collection uploaded successfully",
		Id:  id,
	})
}

// UpdateCollectionsImage
// @Description Create collections
// @Tags posts
// @Accept json
// @Produce json
// @Param img formData file true "image of collections"
// @Param id formData string true "collections identifier"
// @Success 200 {object} domain.OK
// @Router /collections/img [PUT]
func (p *Posts) UpdateCollectionsImage(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	id := c.FormValue("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing 'id' field",
		})
	}

	if err := p.Services.UploadCollectionsImage(c.Request().Context(), id, file); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.Error{
		Msg: "upload image of collection successfully",
	})
}

// GetSlicesOfCollections
// @Description Create collections
// @Tags posts
// @Accept json
// @Produce json
// @Param position query string true "position of collections"
// @Param limit query string true "limit of cards banner slices"
// @Success 200 {object} domain.Collections
// @Router /collections [GET]
func (p *Posts) GetSlicesOfCollections(c echo.Context) error {
	position := c.QueryParam("position")
	limit := c.QueryParam("limit")
	if position == "" || limit == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing 'position' or 'limit' field",
		})
	}

	_limit, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}

	slices, err := p.Services.SlicesOfCollections(c.Request().Context(), position, _limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, *slices)
}
