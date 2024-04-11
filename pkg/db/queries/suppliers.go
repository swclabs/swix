package queries

const (
	InsertIntoSuppliers string = `
		INSERT INTO suppliers (name, phone_number, email)
		VALUES (?, ?, ?);
	`

	InsertIntoSuppliersAddress string = `
		INSERT INTO suppliers_address (suppliers_id, address_uuid)
		VALUES (?, ?);
	`
)
