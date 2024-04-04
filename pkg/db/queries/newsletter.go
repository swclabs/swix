package queries

const (
	InsertIntoNewsletter string = `
		INSERT INTO newsletter (type, title, subtitle, description, image, textcolor)
		VALUES (?, ?, ?, ?, ?, ?);
	`
)
