package addresses

const (
	insertIntoAddresses = `
		INSERT INTO addresses (street, ward, district, city, user_id) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`
	selectAddressesByID = `
		SELECT * FROM addresses WHERE id = $1
	`

	selectAddressesByUserID = `
		SELECT * FROM addresses WHERE user_id = $1
	`
)
