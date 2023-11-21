package queries

const (
	INSERT_INTO_USERS string = `INSERT INTO users (email, phone_number, first_name, last_name, image) 
							    VALUES (?,?,?,?,?)`

	SELECT_USER_INFO string = `SELECT users.email, phone_number, first_name, last_name, image, username, role
					   		   FROM users JOIN accounts ON users.email = accounts.email
					   		   WHERE users.email = ?;`
)
