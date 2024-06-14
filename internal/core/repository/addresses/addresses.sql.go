package addresses

const (
	insertIntoAddresses = `
		INSERT INTO addresses (street, ward, district, city, uuid) 
		VALUES ($1, $2, $3, $4, $5)
	`
)
