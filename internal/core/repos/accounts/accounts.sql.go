package accounts

const (
	insertIntoAccounts string = `
		INSERT INTO accounts (username, role, email, password, created_at, type) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (email) 
		DO UPDATE SET email = accounts.email
		RETURNING id;
	`

	updateAccountsUsername string = `
		UPDATE accounts
		SET username = $1
		WHERE email = $2;
	`

	updateAccountsPassword string = `
		UPDATE accounts
		SET password = $1
		WHERE email = $2;
	`

	updateAccountsRole string = `
		UPDATE accounts
		SET role = $1
		WHERE email = $2;
	`

	selectByEmail string = `
		SELECT * FROM accounts
		WHERE email = $1	
	`
)
