package products

const (
	insertIntoProducts string = `
		INSERT INTO products (image, price, name, description, supplier_id, category_id, status, specs, shop_image)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, '')
		RETURNING id;
	`

	updateProductImage string = `
		UPDATE products
		SET image = $1
		WHERE id = $2;
	`

	updateShopImage string = `
		UPDATE products
		SET shop_image = CASE
						WHEN shop_image IS NULL OR shop_image = '' THEN $1
						ELSE shop_image || ',' || $1
					END
		WHERE id = $2;
	`

	selectLimit string = `
		SELECT *
		FROM products
		LIMIT $1
		OFFSET $2;
	`

	selectByID string = `
		SELECT *
		FROM products
		WHERE id = $1;
	`

	deleteByID = `
		DELETE FROM products
		WHERE id = $1;
	`

	updateByID = `
		UPDATE products
		SET 
			name = CASE
						WHEN $1 <> '' THEN $1
						ELSE name 
					END,
			price = CASE
						WHEN $2 <> '' THEN $2
						ELSE price
					END,
			description = CASE
							WHEN $3 <> '' THEN $3
							ELSE description
						END,
			supplier_id = CASE
							WHEN $4 <> 0 THEN $4
							ELSE supplier_id
						END,
			category_id = CASE
							WHEN $5 <> 0 THEN $5
							ELSE category_id
						END,
			status = CASE
						WHEN $6 <> '' THEN $6
						ELSE status
					END,
			specs = CASE
						WHEN $7 <> '' THEN $7::jsonb
						ELSE specs
					END,
			created = now() AT TIME ZONE 'utc'
		WHERE id = $8;
	`
	searchByKeyword = `
		SELECT *
		FROM products
		WHERE
			name ILIKE '%' || $1 || '%' or 
			description ILIKE '%' || $1 || '%';
	
	`

	selectByCategory = `
		SELECT 
			products.id,
			image, price, products.description, specs,
			products.name as name, 
			categories.name as category_name,
			rating
		FROM 
			products JOIN categories
			ON products.category_id = categories.id
		WHERE categories.name = $1
		OFFSET $2;
	`

	updateRating = `
		UPDATE products SET rating = $1 WHERE id = $2;
	`
)
