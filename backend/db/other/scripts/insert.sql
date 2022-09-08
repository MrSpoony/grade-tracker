USE `gradetracker`;

INSERT INTO `gradetracker`.`tabUser` (`firstname`, `lastname`, `username`, `email`, `password`) VALUES
	("Kimi", "Löffel", "MrSpoony", "mr.spoony@gmail.com", "thisisapassword"),
    ("Mattia", "Gisiger", "maettu999", "maettu999@gmail.com", "apassword"),
    ("Levin", "Zehnder", "levin.999", "levin@gmail.com", "levinspassword")
;

INSERT INTO `gradetracker`.`tabSchool` (`schoolname`) VALUES
	("gibb Berufsfachschule Bern"),
    ("gibb Berufsmaturitätsschule Bern")
;

INSERT INTO `gradetracker`.`tabClass` (`classname`, `school_id`) VALUES
	("INF2021g", 1),
    ("BM1TALS-L2021a", 2)
;

INSERT INTO `gradetracker`.`tabSubject` (`subject`) VALUES
	("Mathematik"),
    ("Deutsch")
;

INSERT INTO `gradetracker`.`tabGrade` (`grade`, `date`, `student_id`, `class_id`, `subject_id`) VALUES
	(5.8, '2022-09-07', 1, 1, 1),
    (4.3, '2022-08-07', 2, 1, 2)
;