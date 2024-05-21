package queries

const (
	InsertIntoCollections = `
		INSERT INTO collections (position, headline, body)
		VALUES (?, ?, ?)
		RETURNING id;
 	`

	SelectCollectionByPosition = `
		SELECT *
		FROM collections
		WHERE position = ?
		LIMIT ?;
	`

	UpdateCollectionImage = `
		UPDATE collections
		SET body = jsonb_set(body, '{image}', to_jsonb(?::text), true)
		WHERE id = ?;
    `
)
