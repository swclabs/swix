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

/***********************************************************************************************/

// HeadlineCollection define in database, used to save and query data from database
type HeadlineCollection struct {
	Id       int64  `json:"id"`
	Created  string `json:"created"`
	Position string `json:"position"`
	Headline string `json:"headline"` // see headline bellows
}

// Headline body of HeadlineCollections used to bind body from json to struct
type Headline struct {
}

// HeadlineType type to accept request and response from users
type HeadlineType struct {
	Id       int64    `json:"id"`
	Created  string   `json:"created"`
	Position string   `json:"position"`
	Headline Headline `json:"headline"`
}
