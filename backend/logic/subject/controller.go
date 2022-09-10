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

func CreateSubject(db *db.DB, subject Subject) (int, error) {
	q := `
INSERT INTO tabSubject (subject) VALUES
(?);
-- sql
`
	getID := `
SELECT MAX(id) FROM tabSubject
-- sql
`
	_, err := db.Query(q, subject.Subject)
	if err != nil {
		return 0, err
	}
	row, err := db.Query(getID)
	if err != nil {
		return 0, err
	}
	var id int
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
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
WHERE id = ?;
-- sql
	`
	_, err := db.Query(q, id)
	return err
}

func GetAllSubjects(db *db.DB) ([]Subject, error) {
	q := `
SELECT id, subject FROM tabSubject
	`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	subjects := []Subject{}
	for rows.Next() {
		subject := Subject{}
		rows.Scan(&subject.ID, &subject.Subject)
		subjects = append(subjects, subject)
	}
	return subjects, nil
}
