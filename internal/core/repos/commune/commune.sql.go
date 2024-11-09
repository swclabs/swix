package commune

const (
	getByDistrictID = `
		SELECT * FROM commune WHERE district_id = $1;
	`
)
