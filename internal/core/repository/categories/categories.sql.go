package categories

const (
	InsertIntoCategory = `
		INSERT INTO categories (name, description) 
		VALUES ($1, $2)
	`

	SelectCategoryLimit string = `
		SELECT *
		FROM categories
		LIMIT $1;
	`
)
