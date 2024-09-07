package comments

const (
	insertIntoComments string = `
		INSERT INTO comments (level, content, product_id, user_id, parent_id, created)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id;	
	`

	selectByID string = `
		SELECT *
		FROM comments
		WHERE id = $1;
	`

	selectCommentsByProductID string = `
		SELECT *
		FROM comments
		WHERE product_id = $1;
	`

	deleteByID string = `
		DELETE FROM comments
		WHERE id = $1;
	`

	updateComments = `
		UPDATE comments
		SET level = CASE
						WHEN $2 <> '' THEN $2
						ELSE level
					END,
			content = CASE
						WHEN $3 <> '' THEN $3
						ELSE content
					END,
			product_id = CASE
						WHEN $4 <> '' THEN $4
						ELSE product_id
					END,
			user_id = CASE
						WHEN $5 <> '' THEN $5
						ELSE user_id
					END
		WHERE id = $1;
		`
)
