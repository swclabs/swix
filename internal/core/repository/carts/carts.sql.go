package carts

const (
	insertItemToCart = `
		INSERT INTO carts (user_id, inventory_id, quantity)
		VALUES ($1, $2, $3);
	`

	selectByUserID = `
		SELECT * FROM carts 
		WHERE user_id = $1
		LIMIT $2;
	`

	deleteItem = `
		DELETE FROM carts 
		WHERE users_id = $1 AND inventory_id = $2
	`
)
