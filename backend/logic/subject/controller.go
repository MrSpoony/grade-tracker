package subject

import (
	"github.com/MrSpoony/grade-tracker/backend/db"
)

func GetSubjectByID(db *db.DB, id int) (*Subject, error) {
	q := `
SELECT id, subject FROM tabSubject
WHERE id = ?
-- sql
`
	row := db.QueryRow(q, id)
	s := Subject{}
	err := row.Scan(&s.ID, &s.Subject)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func CreateSubject(db *db.DB, subject Subject) error {
	q := `
INSERT INTO tabSubject (subject) VALUES
(?)
-- sql
`
	_, err := db.Query(q, subject.Subject)
	return err
}

func UpdateSubject(db *db.DB, subject Subject) error {
	q := `
UPDATE tabSubject
SET subject = ? 
WHERE id = ?
-- sql
`
	_, err := db.Query(q, subject.Subject, subject.ID)
	return err
}

func DeleteSubjectByID(db *db.DB, id int) error {
	q := `
DELETE FROM tabSubject
WHERE id = ?
--sql
	`
	_, err := db.Query(q, id)
	return err
}
