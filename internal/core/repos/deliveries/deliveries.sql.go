package deliveries

const (
	insert = `
		INSERT INTO deliveries (user_id, address_id, status, method, note, sent_date )
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	selectByID = `
		SELECT * FROM deliveries WHERE id = $1
	`

	selectByUserID = ` 
		SELECT * FROM deliveries WHERE user_id = $1
	`
)
