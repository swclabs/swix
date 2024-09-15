package suppliers

const (
	insertIntoSuppliers string = `
		INSERT INTO suppliers (name, email)
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

	// Test
	updateSuppliers string = `
		UPDATE suppliers (name, email)
		VALUES ($1, $2);
	`
)
