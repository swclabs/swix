package warehouse

const (
	insertIntoWarehouse string = `
		INSERT INTO warehouse (product_id, model, price, specs, available, currency_code)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	getAvailableProducts string = `
		SELECT *
		FROM warehouse
		WHERE 
			product_id = $1 AND 
			specs->>'ram' = $2 AND 
			specs->>'ssd' = $3 AND 
			specs->>'color' = $4;
	`

	getById string = `
		SELECT * FROM warehouse where id = $1;
	`
)
