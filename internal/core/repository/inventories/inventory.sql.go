package inventories

const (
	insertIntoInventory string = `
		INSERT INTO inventories (product_id, price, specs, available, currency_code, status)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	getAvailableProducts string = `
		SELECT *
		FROM inventories
		WHERE 
			product_id = $1 AND 
			specs->>'ram' = $2 AND 
			specs->>'ssd' = $3 AND 
			specs->>'color' = $4;
	`

	getByID string = `
		SELECT * FROM inventories WHERE id = $1;
	`

	getByProductID = `
		SELECT * FROM inventories WHERE product_id = $1;
	`

	getProductsLimit = `
		SELECT * FROM inventories LIMIT $1 OFFSET $2;
	`

	deleteInventorybyID = `
		DELETE FROM inventories WHERE id = $1;
	`

	uploadInventoryImage = `
		UPDATE inventories
		SET specs = jsonb_set(
			specs,
			'{image}',
			CASE
				WHEN specs->>'image' IS NULL OR specs->>'image' = 'null' THEN
					to_jsonb(array[$1]::text[])
				ELSE
					(specs->'image')::jsonb || to_jsonb($1::text)
			END,
			true
		)
		WHERE id = $2;
	`

	update = `
		UPDATE inventories
		SET product_id = CASE
							WHEN $2 <> 0 THEN $2
							ELSE product_id
						END,
			status = CASE 
						WHEN $3 <> '' THEN $3 
						ELSE status 
					END,
			price = CASE
						WHEN $4 <> 0 THEN $4
						ELSE price
					END,
			currency_code = CASE
								WHEN $5 <> '' THEN $5
								ELSE currency_code
							END,
			specs = CASE
						WHEN $6 <> '' THEN $6
						ELSE specs
					END,
			available = CASE
							WHEN $7 <> '' THEN $7
							ELSE available
						END
		WHERE id = $1;
	`
)
