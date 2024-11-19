package comments

// type Comment struct {
// 	ID          int64     `json:"id" db:"id"`
// 	Content     string    `json:"content" db:"content"`
// 	UserID      int64     `json:"user_id" db:"user_id"`
// 	ProductID   int64     `json:"product_id" db:"product_id"`
// 	InventoryID int64     `json:"inventory_id" db:"inventory_id"`
// 	StarID      int64     `json:"star_id" db:"star_id"`
// 	Created     time.Time `json:"created" db:"created"`
// }

const (
	insert = `
		INSERT INTO comments (content, user_id, product_id, inventory_id, star_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	getByID = `
		SELECT * FROM comments
		WHERE id = $1
	`

	deleteByID = `
		DELETE FROM comments
		WHERE id = $1
	`

	getByProductID = `
		SELECT * FROM comments
		WHERE product_id = $1
	`

	getModelByProductID = `
		SELECT 
			comment_id as id, user_id, email, first_name, name as product_name,
			last_name, rating, content, specs, color, comment_created as created
		FROM users
		JOIN (
			SELECT 
				content, rating, comments.id as comment_id, 
				comments.created as comment_created, user_id, inventory_id, inventories.specs,
				inventories.color, products.name
			FROM comments 
			JOIN products ON products.id = product_id
			JOIN inventories ON inventories.id = inventory_id
			WHERE products.id = $1
		) ON users.id = user_id;
	`
)
