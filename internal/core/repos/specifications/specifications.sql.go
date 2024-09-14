package specifications

const (
	getByInventoryID = `
		SELECT *
		FROM specifications
		WHERE inventory_id = $1;
	`

	insert = `
		INSERT INTO specifications (inventory_id, content)
		VALUES ($1, $2)
	`

	getByID = `
		SELECT *
		FROM specifications
		WHERE id = $1;
	`
)
