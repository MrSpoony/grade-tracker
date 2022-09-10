package school

import "github.com/MrSpoony/grade-tracker/backend/db"

func GetSchoolByID(db *db.DB, id int) (*School, error) {
	q := `
SELECT id, schoolname FROM tabSchool
WHERE id = ?
-- sql
`
	row := db.QueryRow(q, id)
	s := School{}
	err := row.Scan(&s.ID, &s.Schoolname)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
