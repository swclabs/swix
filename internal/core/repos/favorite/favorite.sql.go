package favorite

// type Favorite struct {
// 	ID          int64 `json:"id" db:"id"`
// 	UserID      int64 `json:"user_id" db:"user_id"`
// 	InventoryID int64 `json:"inventory_id" db:"inventory_id"`
// }

const (
	insert = `
		INSERT INTO favorite (user_id, inventory_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, inventory_id)
		DO NOTHING;
	`

	delete = `
		DELETE FROM favorite
		WHERE user_id = $1 AND inventory_id = $2;
	`

	getByInventoryID = `
		SELECT *
		FROM favorite
		WHERE inventory_id = $1 AND user_id = $2;
	`

	getByUserID = `	
		SELECT *
		FROM favorite
		WHERE user_id = $1;
	`
)
