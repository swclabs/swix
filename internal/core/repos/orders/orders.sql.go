package orders

const (
	insertOrder = `
		INSERT INTO orders (uuid, user_id, status, total_amount, delivery_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	insertProductToOrder = `
		INSERT INTO product_in_order (order_id, inventory_id, quantity, currency_code, total_amount)
		VALUES ($1, $2, $3, $4, $5);
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

	getByUUID = `
		SELECT * FROM orders WHERE uuid = $1;
	`

	getByOrderCode = `
		SELECT uuid, time, user_id, total_amount, quantity, currency_code, color, inventory_image, name, category_id, content FROM (
			SELECT uuid, time, user_id, total_amount, specs_id, quantity, currency_code, color, inventory_image, name, category_id FROM (
				SELECT uuid, time, user_id, total_amount, specs_id, quantity, inventories.currency_code, color, image as inventory_image, product_id FROM (
					SELECT product_in_order.order_id, uuid, time, user_id, delivery_id, product_in_order.total_amount, status, inventory_id, specs_id, quantity, currency_code FROM (
						SELECT id as o_id, uuid, time, user_id, delivery_id, total_amount, status 
						FROM orders where orders.uuid = $1
					) JOIN product_in_order ON product_in_order.order_id = o_id
				) JOIN inventories ON inventories.id = inventory_id
			) JOIN products ON products.id = product_id
		) JOIN specifications ON specifications.id = specs_id;
	`
)
