package suppliers

const (
	insertIntoSuppliers string = `
		INSERT INTO suppliers (name, email)
		VALUES ($1, $2);
	`

	insertIntoSuppliersAddress string = `
		INSERT INTO suppliers_address (suppliers_id, address_uuid)
		VALUES ($1, $2);
	`

	selectSupplierByEmailLimit string = `
		SELECT * 
		FROM suppliers
		LIMIT $1;
	`

	selectByEmail string = `
		SELECT *
		FROM suppliers
		WHERE email = $1
		LIMIT 1;
	`
)
