package queries

const (
	InsertIntoAddresses = `
		INSERT INTO addresses (street, ward, district, city, uuid) 
		VALUES (?, ?, ?, ?, ?)
	`
)
