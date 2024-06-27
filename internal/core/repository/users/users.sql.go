package users

const (
	insertIntoUsers string = `
		INSERT INTO users (email, phone_number, first_name, last_name, image) 
		VALUES ($1,$2,$3,$4,$5)
	`

	selectUserInfo string = `
		SELECT users.id, users.email, phone_number, first_name, last_name, image, username, role
		FROM users 
		JOIN accounts ON users.email = accounts.email
		WHERE users.email = $1;
	`

	updateUsersFirstname string = `
		UPDATE users
		SET first_name = $1
		WHERE email = $2;
	`

	updateUsersImage string = `
		UPDATE users
		SET image = $1
		WHERE email = $2;
	`

	updateUsersLastname string = `
		UPDATE users
		SET last_name = $1
		WHERE email = $2;
	`

	updateUsersPhoneNumber string = `
		UPDATE users
		SET phone_number = $1
		WHERE email = $2;
	`

	insertUsersConflict string = `
		INSERT INTO users (email, phone_number, first_name, last_name, image) 
		VALUES ($1, $2, $3, $4, $5) 
		ON CONFLICT (email)
		DO
			UPDATE 
			SET first_name = EXCLUDED.first_name, 
				last_name = EXCLUDED.last_name,
				image = EXCLUDED.image;
	`
	selectByEmail string = `
		SELECT *
		FROM users
		WHERE email = $1;
	
	`

	selectByPhone string = `
		SELECT *
		FROM users
		WHERE phone_number = $1;
	`
)
