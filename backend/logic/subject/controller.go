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
