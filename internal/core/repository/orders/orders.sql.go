package orders

const (
	insertOrder = `
		INSERT INTO orders (uuid, user_id, status, total_amount, delivery_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	insertProductToOrder = `
		INSERT INTO product_in_order (order_id, inventory_id, quantity, currency_code, total_amount, specs_id)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	getOrder = `
		SELECT *
		FROM orders
		WHERE user_id = $1
		ORDER BY id DESC
		LIMIT $2;
	`

	getProductByOrderID = `
		SELECT * FROM product_in_order WHERE order_id = $1 ORDER BY id ASC;
	`
)
