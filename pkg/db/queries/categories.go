package queries

const (
	InsertIntoCategory = `
		INSERT INTO categories (name, description) 
		VALUES (?, ?)
	`
)
