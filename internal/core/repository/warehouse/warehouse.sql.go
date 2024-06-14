package warehouse

const (
	InsertIntoWarehouse string = `
		INSERT INTO warehouse (product_id, model, price, specs, available, currency_code)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	GetAvailableProducts string = `
		SELECT *
		FROM warehouse
		WHERE 
			product_id = $1 AND 
			specs->>'ram' = $2 AND 
			specs->>'ssd' = $3 AND 
			specs->>'color' = $4;
	`
)
