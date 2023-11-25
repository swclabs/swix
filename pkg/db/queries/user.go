package queries

const (
	InsertIntoUsers string = `
		INSERT INTO users (email, phone_number, first_name, last_name, image) 
		VALUES (?,?,?,?,?)
	`

	SelectUserInfo string = `
		SELECT users.id, users.email, phone_number, first_name, last_name, image, username, role
		FROM users 
		JOIN accounts ON users.email = accounts.email
		WHERE users.email = ?;
	`

	UpdateUsersFirstname string = `
		UPDATE users
		SET first_name = ?
		WHERE email = ?;
	`

	UpdateUsersImage string = `
		UPDATE users
		SET image = ?
		WHERE email = ?;
	`

	UpdateUsersLastname string = `
		UPDATE users
		SET last_name = ?
		WHERE email = ?;
	`

	UpdateUsersPhoneNumber string = `
		UPDATE users
		SET phone_number = ?
		WHERE email = ?;
	`
)
