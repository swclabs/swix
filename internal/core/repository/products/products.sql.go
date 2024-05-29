package products

const (
	InsertIntoProducts string = `
		INSERT INTO products (image, price, name, description, supplier_id, category_id, status, spec)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	UpdateProductImage string = `
		UPDATE products
		SET image = image || ',' || ?
		WHERE id = ?;
	`
)
