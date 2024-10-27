package stars

const (
	insertStar = `
		INSERT INTO stars (product_id, user_id) 
		VALUES ($1, $2) ON CONFLICT (product_id, user_id) 
		DO NOTHING;
	`
)
