package dtos

// CollectionUpload request, response
type CollectionUpload struct {
	Msg string `json:"msg"`
	ID  int64  `json:"id"`
}

// Headline request, response
type Headline struct {
	Content []string `json:"content"`
}

// HeadlineContent request, response
type HeadlineContent struct {
	Content string `json:"content"`
}

// Message request, response
type Message struct {
	Position string   `json:"position" validate:"required"`
	Content  []string `json:"content" validate:"required"`
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

type Comment struct {
	// Position string   `json:"position" validate:"required"`
	Content []string `json:"content" validate:"required"`
	Name    string   `json:"name" validate:"required"`
	Level   int64    `json:"level" validate:"required"`
	Rating  int64    `json:"rating" validate:"required"`
	Like    int64    `json:"like" validate:"required"`
	Dislike int64    `json:"dislike" validate:"required"`
}
