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

// CardArticle used to bind from json mapping to structure
type CardArticle struct {
	Category string        `json:"category" validate:"required"`
	Title    string        `json:"title" validate:"required"`
	Src      string        `json:"src"`
	Content  []CardContent `json:"content"`
}

// Article request, response
type Article struct {
	Headline string        `json:"headline" validate:"required"`
	Cards    []CardArticle `json:"cards" validate:"required"`
}

// UploadArticle request, response
type UploadArticle struct {
	Position string `json:"position" validate:"required"`
	Article
}

type News struct {
	Header string        `json:"header" validate:"required"`
	Cards  []CardArticle `json:"cards" validate:"required"`
}

type NewsDTO struct {
	Header   string        `json:"header" validate:"required"`
	Cards    []CardArticle `json:"cards" validate:"required"`
	Category string        `json:"category" validate:"required"`
}

type Comment struct {
	// Position string   `json:"position" validate:"required"`
	ID        int64    `json:"id" validate:"required"`
	Content   []string `json:"content" validate:"required"`
	Username  string   `json:"username" validate:"required"`
	Level     int64    `json:"level" validate:"required"` // 0: parent, 1: child
	ParentID  int64    `json:"parent_id"`
	Rating    int64    `json:"rating" validate:"required"`
	Liked     int64    `json:"like" validate:"required"`
	Disliked  int64    `json:"dislike" validate:"required"`
	UserID    int64    `json:"user_id" db:"user_id"`
	ProductID int64    `json:"product_id" db:"product_id"`
}
