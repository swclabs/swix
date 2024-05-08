package queries

const (
	InsertIntoProducts string = `
		INSERT INTO products (image, price, name, description, available, supplier_id, category_id, status, spec)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	UpdateProductImage string = `
		UPDATE products
		SET image = image || ',' || ?
		WHERE id = ?;
	
	`
)
