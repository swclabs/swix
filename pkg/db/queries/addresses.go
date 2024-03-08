package queries

const (
	InsertIntoAddresses = `
		INSERT INTO addresses (user_id, supplier_id, street, ward, district, city) 
		VALUES (?, ?, ?, ?, ?, ?)
	`
)
