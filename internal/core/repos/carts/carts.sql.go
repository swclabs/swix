package carts

const (
	insertItemToCart = `
		INSERT INTO carts (user_id, inventory_id, quantity, spec_id)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (inventory_id, user_id, spec_id)
		DO UPDATE SET quantity = carts.quantity + EXCLUDED.quantity;
	`

	selectByUserID = `
		SELECT * FROM carts 
		WHERE user_id = $1
		LIMIT $2;
	`

	deleteItem = `
		DELETE FROM carts 
		WHERE id = $1;
	`
)
