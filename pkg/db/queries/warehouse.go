package queries

const (
	InsertIntoWarehouse string = `
		INSERT INTO warehouse (product_id, model, price, specs, available)
		VALUES (?, ?, ?, ?, ?);
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
