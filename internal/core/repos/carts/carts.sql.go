package carts

const (
	insertItemToCart = `
		INSERT INTO carts (user_id, inventory_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (inventory_id, user_id)
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

	getCartInfo = `
		SELECT
			cart_id,
			inventory_id,
			quantity,
			product_id,
			inventory_price,
			currency_code,
			color,
			inventory_specs,
			products.name AS name,
			image AS inventory_image,
			categories.name AS category_name
		FROM
		(
			SELECT
				color,
				cart_id,
				quantity,
				product_id,
				inventory_id,
				currency_code,
				price AS inventory_price,
				image AS inventory_image,
				inventories.specs AS inventory_specs
			FROM
			(
				SELECT
					id AS cart_id,
					inventory_id,
					quantity
				FROM
					carts
				WHERE
					user_id = $1
			)
			JOIN inventories ON inventory_id = inventories.id
		)
		JOIN products ON products.id = product_id
		JOIN categories ON categories.id = products.category_id;
	`
)
