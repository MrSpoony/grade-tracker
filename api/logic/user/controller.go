package user

import (
	"github.com/MrSpoony/grade-tracker/api/db"
)

func GetUserByUsername(db *db.DB, username string) (*User, error) {
	q := `
SELECT id, firstname, lastname, username, email, password FROM tabUser
WHERE username = ?
-- sql
`
	row := db.QueryRow(q, username)
	u := User{}
	err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func StoreNewUser(db *db.DB, user User) error {
	q := `
INSERT INTO tabUser (firstname, lastname, username, email, password) VALUES
(?, ?, ?, ?, ?)
-- sql
`
	_, err := db.Query(q, user.Firstname, user.Lastname, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
