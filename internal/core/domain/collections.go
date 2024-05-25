package domain

// Collection use to query data from database, define in database, used to save a CollectionBody
type Collection struct {
	Id       int64  `json:"id"`
	Created  string `json:"created"`
	Position string `json:"position"` // Example: mac#1, mac#2
	Headline string `json:"headline"` // Ex: Get to know Mac
	Body     string `json:"body"`
}

/***********************************************************************************************/

// LocalizationCollection is a structure define in Database used to save a Localization json file
type LocalizationCollection struct {
	Id           int64  `json:"id"`
	Localization string `json:"localization"` // see Localization bellows
}

// Localization is a struct to bind body of LocalizationCollection to structure type
type Localization struct {
}

// LocalizationType is a type used to request and responses
type LocalizationType struct {
	Id           int64        `json:"id"`
	Localization Localization `json:"localization"`
}

/***********************************************************************************************/

// CollectionBody used to bind from json mapping to structure
type CollectionBody struct {
	Title       string `json:"title" validate:"required"`
	SubTitle    string `json:"subtitle" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"`
	TextColor   string `json:"text_color" validate:"required"`
}

// CollectionsBody is a body of CardBannerSlice
type CollectionsBody struct {
	CollectionBody
	Id      int64  `json:"id"`
	Created string `json:"created"`
}

// CollectionType is a type use to accept request and response
type CollectionType struct {
	Id       int64          `json:"id"`
	Created  string         `json:"created"`
	Position string         `json:"position" validate:"required"`
	Headline string         `json:"headline" validate:"required"`
	Body     CollectionBody `json:"body" validate:"required"`
}

// Collections is a type use to accept request and response
type Collections struct {
	Position   string            `json:"position"`
	Headline   string            `json:"headline"`
	CardBanner []CollectionsBody `json:"card_banner"`
}

type CollectionUploadRes struct {
	Msg string `json:"msg"`
	Id  int64  `json:"id"`
}

/*** Swagger ***/

type CollectionTypeSwagger struct {
	Position string                `json:"position" validate:"required"`
	Headline string                `json:"headline" validate:"required"`
	Body     CollectionBodySwagger `json:"body" validate:"required"`
}

type CollectionBodySwagger struct {
	Title       string `json:"title" validate:"required"`
	SubTitle    string `json:"subtitle" validate:"required"`
	Description string `json:"description" validate:"required"`
	TextColor   string `json:"text_color" validate:"required"`
}

/***********************************************************************************************/

type HeadlineBannerBody struct {
	Headline string `json:"headline" validate:"required"`
}

type HeadlineBannerSliceBody struct {
	HeadlineBannerBody
	Id      int64  `json:"id"`
	Created string `json:"created"`
}

// HeadlineBannerType user body request & response
type HeadlineBannerType struct {
	Position string             `json:"position" validate:"required"`
	Created  string             `json:"created"`
	Body     HeadlineBannerBody `json:"body" validate:"required"`
}

// HeadlineBannerTypeSwagger used to generate swagger documents
type HeadlineBannerTypeSwagger struct {
	Position string             `json:"position" validate:"required"`
	Body     HeadlineBannerBody `json:"body" validate:"required"`
}

// HeadlineBannerSlice response slices
type HeadlineBannerSlice struct {
	Position  string                    `json:"position"`
	Headlines []HeadlineBannerSliceBody `json:"headlines"`
}
