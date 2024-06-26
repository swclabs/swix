package carts

const (
	insertItemToCart = `
		INSERT INTO carts (user_id, warehouse_id, quantity)
		VALUES ($1, $2, $3);
	`

	selectByUserId = `
		SELECT * FROM carts 
		WHERE user_id = $1
		LIMIT $2;
	`

	deleteItem = `
		DELETE FROM carts 
		WHERE users_id = $1 AND warehouse_id = $2
	`
)
