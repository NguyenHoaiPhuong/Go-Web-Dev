package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	pgUsername = "postgres"
	pgPassword = "postgres"
	pgDBName   = "akagi"
	pgPort     = 5432
	pgHost     = "127.0.0.1"
)

func initPostgres() *sqlx.DB {
	dbConfig := fmt.Sprintf("user=%s dbname=%s host=%s port=%d sslmode=disable password=%s",
		pgUsername, pgDBName, pgHost,
		pgPort, pgPassword)
	log.Printf("Init db with these param %v", dbConfig)

	return sqlx.MustConnect("postgres", dbConfig)
}

type (
	// User :
	User struct {
		ID      int     `json:"id" db:"id"`
		Name    string  `json:"name" db:"name"`
		Courses Courses `json:"courses" db:"courses"`
	}

	// Course :
	Course struct {
		ID       int      `json:"id" db:"id"`
		Title    string   `json:"title" db:"title"`
		Teachers Teachers `json:"teachers" db:"teachers"`
	}

	// Courses :
	Courses []Course

	// UserCourse :
	UserCourse struct {
		UserID   int `json:"user_id" db:"user_id"`
		CourseID int `json:"course_id" db:"course_id"`
		Role     int `json:"role" db:"role"`
	}

	// Teacher :
	Teacher struct {
		ID   int    `json:"id" db:"id"`
		Name string `json:"name" db:"name"`
	}

	// Teachers :
	Teachers []Teacher

	// CourseTeacher :
	CourseTeacher struct {
		CourseID  int `json:"course_id" db:"course_id"`
		TeacherID int `json:"teacher_id" db:"teacher_id"`
	}
)

// Scan :
func (c *Courses) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return errors.New("Incompatible type for Courses, byte or string")
	}

	return json.Unmarshal(source, c)
}

// Scan :
func (t *Teachers) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return errors.New("Incompatible type for Teachers, byte or string")
	}

	return json.Unmarshal(source, t)
}

func main() {
	db := initPostgres()
	if err := db.Ping(); err != nil {
		panic(err)
	}

	creteTables(db)
	insertUsers(db)
	insertTeachers(db)
	insertCourses(db)
	insertCourseTeachers(db)
	insertUserCourses(db)

	query := `
	SELECT
		u.*,
		JSON_AGG(
			JSON_BUILD_OBJECT(
				'id', c.id,
				'title', c.title
			)
		) AS courses
	FROM users u, courses c, user_courses uc, course_teachers ct, teachers t
	WHERE
		uc.user_id = u.id 
		AND uc.course_id = c.id
		AND ct.teacher_id = t.id
		AND ct.course_id = c.id
	GROUP BY u.id`

	var users []*User
	err := db.Select(&users, query)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

func creteTables(db *sqlx.DB) {
	query := `
	CREATE TABLE users (
		id	SERIAL PRIMARY KEY,
		name	TEXT UNIQUE
	);`
	res, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	log.Println(res)

	query = `
	CREATE TABLE courses (
		id	SERIAL PRIMARY KEY,
		title	TEXT UNIQUE
	);`
	res, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	log.Println(res)

	query = `
	CREATE TABLE teachers (
		id	SERIAL PRIMARY KEY,
		name	TEXT UNIQUE
	);`
	res, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	log.Println(res)

	query = `
	CREATE TABLE user_courses (
		user_id	INTEGER REFERENCES users(id),
		course_id	INTEGER REFERENCES courses(id),
		role	INTEGER,
		PRIMARY KEY(user_id,course_id)
	);`
	res, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	log.Println(res)

	query = `
	CREATE TABLE course_teachers (
		teacher_id	INTEGER REFERENCES teachers(id),
		course_id	INTEGER REFERENCES courses(id),
		PRIMARY KEY(teacher_id,course_id)
	);`
	res, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	log.Println(res)
}

func insertUsers(db *sqlx.DB) {
	query := `
	INSERT INTO users(name)
	VALUES ('akagi'), ('yushin'), ('mogami');`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func insertTeachers(db *sqlx.DB) {
	query := `
	INSERT INTO teachers(name)
	VALUES ('A'), ('B'), ('C');`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func insertCourses(db *sqlx.DB) {
	query := `
	INSERT INTO courses(title)
	VALUES ('golang'), ('python'), ('C/C++');`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func insertCourseTeachers(db *sqlx.DB) {
	query := `
	INSERT INTO course_teachers(teacher_id, course_id)
	VALUES (1, 1), (1, 2), (2, 2), (2, 3), (3, 3), (3, 1);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func insertUserCourses(db *sqlx.DB) {
	query := `
	INSERT INTO user_courses(user_id, course_id, role)
	VALUES (1, 1, 1), (1, 2, 1), (2, 2, 0), (2, 3, 0), (3, 3, 0), (3, 1, 0);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
