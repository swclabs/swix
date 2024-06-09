package warehouse

const (
	InsertIntoWarehouse string = `
		INSERT INTO warehouse (product_id, model, price, specs, available, currency_code)
		VALUES (?, ?, ?, ?, ?, ?);
	`

	GetAvailableProducts string = `
		SELECT *
		FROM warehouse
		WHERE 
			product_id = ? AND 
			specs->>'ram' = ? AND 
			specs->>'ssd' = ? AND 
			specs->>'color' = ?;
	`
)
