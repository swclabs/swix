package inventories

const (
	insertIntoInventory string = `
		INSERT INTO inventories (product_id, model, price, specs, available, currency_code, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
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
)
