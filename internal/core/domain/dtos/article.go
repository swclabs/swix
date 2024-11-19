package dtos

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
	Star        int    `json:"star" validate:"required"`
	Content     string `json:"content" validate:"required"`
	UserID      int64  `json:"user_id" validate:"required"`
	ProductID   int64  `json:"product_id" validate:"required"`
	InventoryID int64  `json:"inventory_id" validate:"required"`
}

type CommentResp struct {
}
