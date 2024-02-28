package queries

const (
	InsertIntoAccounts string = `
		INSERT INTO accounts (username, role, email, password, created_at, type) 
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT (email) 
		DO NOTHING;
	`

	UpdateAccountsUsername string = `
		UPDATE accounts
		SET username = ?
		WHERE email = ?;
	`

	UpdateAccountsPassword string = `
		UPDATE accounts
		SET password = ?
		WHERE email = ?;
	`

	UpdateAccountsRole string = `
		UPDATE accounts
		SET role = ?
		WHERE email = ?;
	`
)
