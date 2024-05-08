package queries

const (
	InsertIntoWarehouse string = `
		INSERT INTO warehouse (product_id, model, ram, ssd, available)
		VALUES (?, ?, ?, ?, ?);
	`

	GetAvailableProducts string = `
		SELECT *
		FROM warehouse
		WHERE product_id = ? AND ram = ? AND ssd = ?;
	`
)
