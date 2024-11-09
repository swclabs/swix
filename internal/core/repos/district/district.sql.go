package district

const (
	getByProvinceID = `	
		SELECT * FROM district WHERE province_id = $1;
	`
)
