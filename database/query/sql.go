package query

var (
	InsertNewUserQuery string = `
		insert into users (username, password) values
		($1, $2)
	`

	GetUserByIDQuery string = `
		select username, password, is_active from users where id = ($1)
	`

	GetUserByUsernameQuery string = `
		select username, password, is_active from users where username = ($1)
	`

	GetAllUsersQuery string = `
		select username, password, is_active from users
	`
)
