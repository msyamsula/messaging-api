package query

var (
	InsertNewUserQuery string = `
		insert into users (username, password) values
		($1, $2)
	`

	GetUserByIDQuery string = `
		select id, username, password, is_active from users where id = ($1)
	`

	GetUserByUsernameQuery string = `
		select id, username, password, is_active from users where username = ($1)
	`

	GetAllUsersQuery string = `
		select id, username, password, is_active from users order by username asc
	`

	SetIsActiveUserQuery string = `
		update users set is_active = ($1) where username = ($2)
	`
)
