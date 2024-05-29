package categories

const (
	InsertIntoCategory = `
		INSERT INTO categories (name, description) 
		VALUES (?, ?)
	`

	SelectCategoryLimit string = `
		SELECT *
		FROM categories
		LIMIT ?;
	`
)
