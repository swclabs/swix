package queries

const (
	InsertIntoUsers string = `
		INSERT INTO users (email, phone_number, first_name, last_name, image) 
		VALUES (?,?,?,?,?)
	`

	SelectUserInfo string = `
		SELECT users.id, users.email, phone_number, first_name, last_name, image, username, role
		FROM users JOIN accounts ON users.email = accounts.email
		WHERE users.email = ?;
	`
)
