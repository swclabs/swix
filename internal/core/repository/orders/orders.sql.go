package orders

const (
	insertOrder = `
		INSERT INTO orders (uuid, user_id, status, total_amount)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	insertProductToOrder = `
		INSERT INTO product_in_order (order_id, inventory_id, quantity, currency_code, total_amount)
		VALUES ($1, $2, $3, $4, $5);
	`

	getOrder = `
		SELECT *
		FROM orders
		ORDER BY id DESC
		WHERE user_id = $1
		LIMIT $2;
	`

	getProductByOrderID = `
		SELECT * FROM product_in_order WHERE order_id = $1 ORDER BY id ASC;
	`
)
