package collections

const (
	InsertIntoCollections = `
		INSERT INTO collections (position, headline, body)
		VALUES ($1, $2, $3)
		RETURNING id;
 	`

	SelectCollectionByPosition = `
		SELECT *
		FROM collections
		WHERE position = $1
		LIMIT $2;
	`

	UpdateCollectionImage = `
		UPDATE collections
		SET body = jsonb_set(body, '{image}', to_jsonb(?::text), true)
		WHERE id = $1;
    `
)
