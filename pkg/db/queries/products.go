package queries

const (
	InsertIntoProducts string = `
		INSERT INTO products (image, price, name, description, available, supplier_id, category_id)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	UpdateProductImage string = `
		UPDATE products
		SET image = image || ',' || ?
		WHERE id = ?;
	
	`
)
