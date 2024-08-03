package dtos

// CollectionUpload request, response
type CollectionUpload struct {
	Msg string `json:"msg"`
	ID  int64  `json:"id"`
}

// HeadlineBannerBody request, response
type HeadlineBannerBody struct {
	Headline string `json:"headline" validate:"required"`
}

// HeadlineBannerSlicesBody request, response
type HeadlineBannerSlicesBody struct {
	HeadlineBannerBody
	ID      int64  `json:"id"`
	Created string `json:"created"`
}

// HeadlineBanner user body request & response
type HeadlineBanner struct {
	Position string             `json:"position" validate:"required"`
	Created  string             `json:"created"`
	Body     HeadlineBannerBody `json:"body" validate:"required"`
}

// HeadlineBanners response slices
type HeadlineBanners struct {
	Position  string                     `json:"position"`
	Headlines []HeadlineBannerSlicesBody `json:"headlines"`
}

// CardContent used to bind from json mapping to structure
type CardContent struct {
	Content string `json:"content" validate:"required"`
	Src     string `json:"src"`
}

// CardInArticle used to bind from json mapping to structure
type CardInArticle struct {
	Category string        `json:"category" validate:"required"`
	Title    string        `json:"title" validate:"required"`
	Src      string        `json:"src"`
	Content  []CardContent `json:"content"`
}

// Article request, response
type Article struct {
	Headline string          `json:"headline" validate:"required"`
	Cards    []CardInArticle `json:"cards" validate:"required"`
}

// UploadArticle request, response
type UploadArticle struct {
	Position string `json:"position" validate:"required"`
	Article
}
