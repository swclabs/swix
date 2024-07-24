package dto

// CollectionBody used to bind from json mapping to structure
type CollectionBody struct {
	Title       string `json:"title" validate:"required"`
	SubTitle    string `json:"subtitle" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"`
	TextColor   string `json:"text_color" validate:"required"`
}

// CollectionSliceBody is a body of CardBannerSlice
type CollectionSliceBody struct {
	CollectionBody
	ID      int64  `json:"id"`
	Created string `json:"created"`
}

// Collection is a type use to accept request and response
type Collection struct {
	ID       int64          `json:"id"`
	Created  string         `json:"created"`
	Position string         `json:"position" validate:"required"`
	Headline string         `json:"headline" validate:"required"`
	Body     CollectionBody `json:"body" validate:"required"`
}

// Collections is a type use to accept request and response
type Collections struct {
	Position   string                `json:"position"`
	Headline   string                `json:"headline"`
	CardBanner []CollectionSliceBody `json:"card_banner"`
}

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
