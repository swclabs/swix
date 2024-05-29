package suppliers

const (
	InsertIntoSuppliers string = `
		INSERT INTO suppliers (name, email)
		VALUES (?, ?);
	`

	InsertIntoSuppliersAddress string = `
		INSERT INTO suppliers_address (suppliers_id, address_uuid)
		VALUES (?, ?);
	`
)
