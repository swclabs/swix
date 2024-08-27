package comments

const (
	insertIntoComments string = `
		INSERT INTO comments (level, content, product_id, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;	
	`

	selectCommentByID string = `
		SELECT *
		FROM comments
		WHERE id = $1;
		`
)
