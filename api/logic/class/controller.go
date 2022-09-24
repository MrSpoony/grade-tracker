package class

import (
	"errors"

	"github.com/MrSpoony/grade-tracker/api/db"
	"github.com/MrSpoony/grade-tracker/api/logic/school"
)

func GetClassByID(db *db.DB, id int) (*Class, error) {
	q := `
SELECT id, classname, school_id FROM tabClass
WHERE id = ?
-- sql
`
	row := db.QueryRow(q, id)
	c := Class{}
	err := row.Scan(&c.ID, &c.Classname, &c.SchoolID)
	if err != nil {
		return nil, err
	}
	school, err := school.GetSchoolByID(db, c.SchoolID)
	if err != nil {
		return nil, err
	}
	c.School = *school
	return &c, nil
}

func GetAllClasses(db *db.DB) ([]Class, error) {
	q := `
SELECT id, classname, school_id FROM tabClass
-- sql
`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	classes := []Class{}
	for rows.Next() {
		c := Class{}
		err := rows.Scan(&c.ID, &c.Classname, &c.SchoolID)
		if err != nil {
			return nil, err
		}
		school, err := school.GetSchoolByID(db, c.SchoolID)
		if err != nil {
			return nil, err
		}
		c.School = *school
		classes = append(classes, c)
	}
	return classes, nil
}

func CreateClass(db *db.DB, class Class) error {
	q := `
INSERT INTO tabClass (classname, school_id) VALUES
(?, ?)
-- sql
`
	s := school.School{}
	if class.School != s {
		// TODO: Handle new creation of school with class
		return errors.New("not implemented yet")
	}
	_, err := db.Query(q, class.Classname, class.SchoolID)
	return err
}
