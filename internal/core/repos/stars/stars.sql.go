package stars

const (
	insertStar = `
		INSERT INTO stars (product_id, user_id, star) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (product_id, user_id) 
		DO UPDATE SET id = -1 WHERE FALSE 
		RETURNING COALESCE(id, -1) AS result_id;
	`
	getByProductID = `
		SELECT * FROM stars WHERE product_id = $1;
	`
)
