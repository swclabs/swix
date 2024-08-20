package categories

const (
	insertIntoCategory = `
		INSERT INTO categories (name, description) 
		VALUES ($1, $2)
	`

	selectCategoryLimit string = `
		SELECT *
		FROM categories
		LIMIT $1;
	`

	selectCategoryByID string = `
		SELECT *
		FROM categories
		WHERE id = $1;
	`

	deleteByID = `
		DELETE FROM categories
		WHERE id = $1;
	`

	updateCategories = `
		UPDATE suppliers
		SET name = CASE
						WHEN $2 <> '' THEN $2
						ELSE name
					END,
			description = CASE
						WHEN $3 <> '' THEN $3
						ELSE description
					END
		WHERE id = $1;
	`
)
