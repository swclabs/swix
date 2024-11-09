package news

const (
	insertIntoNews = `
		INSERT INTO news (category, header, body)
		VALUES ($1, $2, $3)
		RETURNING id;
 	`

	selectByCategory = `
		SELECT *
		FROM news
		WHERE category = $1
		LIMIT $2;
	`

	updateNewsImage = `
		UPDATE news
		SET body = jsonb_set(body, '{image}', to_jsonb($1::text), true)
		WHERE id = $2;
    `
)
