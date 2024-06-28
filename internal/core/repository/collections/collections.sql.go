package collections

const (
	insertIntoCollections = `
		INSERT INTO collections (position, headline, body)
		VALUES ($1, $2, $3)
		RETURNING id;
 	`

	selectCollectionByPosition = `
		SELECT *
		FROM collections
		WHERE position = $1
		LIMIT $2;
	`

	updateCollectionImage = `
		UPDATE collections
		SET body = jsonb_set(body, '{image}', to_jsonb(?::text), true)
		WHERE id = $1;
    `
)
