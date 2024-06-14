package accounts

const (
	InsertIntoAccounts string = `
		INSERT INTO accounts (username, role, email, password, created_at, type) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (email) 
		DO NOTHING;
	`

	UpdateAccountsUsername string = `
		UPDATE accounts
		SET username = $1
		WHERE email = $2;
	`

	UpdateAccountsPassword string = `
		UPDATE accounts
		SET password = $1
		WHERE email = $2;
	`

	UpdateAccountsRole string = `
		UPDATE accounts
		SET role = $1
		WHERE email = $2;
	`

	SelectByEmail string = `
		SELECT * FROM accounts
		WHERE email = $1	
	`
)
