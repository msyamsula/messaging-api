package service

import (
	"database/sql"
	"fmt"

	"github.com/msyamsula/messaging-api/database"
	"github.com/msyamsula/messaging-api/database/query"
)

type DbObject struct {
	D *sql.DB
}

func New(c database.DbConfig) (*DbObject, error) {

	dbo := &DbObject{}
	d, err := connect(c)
	if err != nil {
		return dbo, err
	}

	dbo.D = d
	return dbo, err
}

func connect(c database.DbConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	return db, err
}

func (d *DbObject) InsertUser(username string, password string) error {
	res, err := d.D.Exec(query.InsertNewUserQuery, username, password)
	fmt.Println(res)
	return err
}

func (d *DbObject) GetUserByID(id string) (database.User, error) {
	row := d.D.QueryRow(query.GetUserByIDQuery, id)
	user := database.User{}
	err := row.Scan(&user.Name, &user.Password, &user.IsActive)

	switch err {
	case sql.ErrNoRows:
		return user, nil
	case nil:
		return user, err
	default:
		return user, err
	}
}

func (d *DbObject) GetUserByUsername(username string) (database.User, error) {
	row := d.D.QueryRow(query.GetUserByUsernameQuery, username)
	user := database.User{}
	err := row.Scan(&user.Name, &user.Password, &user.IsActive)

	switch err {
	case sql.ErrNoRows:
		return user, nil
	case nil:
		return user, err
	default:
		return user, err
	}
}

func (d *DbObject) GetAllUser() ([]database.User, error) {
	users := []database.User{}
	var err error
	rows, err := d.D.Query(query.GetAllUsersQuery)
	defer rows.Close()

	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := database.User{}
		err := rows.Scan(&user.Name, &user.Password, &user.IsActive)
		if err != nil {
			continue
		}
		users = append(users, user)
	}

	return users, err

}
