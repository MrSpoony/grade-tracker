package class

import "github.com/MrSpoony/grade-tracker/backend/logic/school"

type Class struct {
	ID        int    `json:"id"`
	Classname string `json:"classname"`
	SchoolID  int    `json:"school_id"`
	School    school.School
}
