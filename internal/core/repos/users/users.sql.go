package users

const (
	insertIntoUsers string = `
		INSERT INTO users (email, phone_number, first_name, last_name, image) 
		VALUES ($1,$2,$3,$4,$5)
		RETURNING id;
	`

	selectUserInfo string = `
		SELECT users.id, users.email, phone_number, first_name, last_name, image, username, role
		FROM users 
		JOIN accounts ON users.email = accounts.email
		WHERE users.email = $1;
	`
	updateInfo string = `
		UPDATE users
		SET first_name = CASE 
							WHEN $2 <> '' THEN $2
							ELSE first_name 
						END,
			last_name = CASE
							WHEN $3 <> '' THEN $3
							ELSE last_name
						END,
			image = CASE
						WHEN $4 <> '' THEN $4
						ELSE image
					END,
			phone_number = CASE
							WHEN $5 <> '' THEN $5
							ELSE phone_number
						END
		WHERE email = $1;
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
	selectByID string = `
		SELECT *
		FROM users
		WHERE id = $1;
	`

	selectByPhone string = `
		SELECT *
		FROM users
		WHERE phone_number = $1;
	`
)
