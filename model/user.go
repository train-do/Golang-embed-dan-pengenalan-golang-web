package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
}

// type User struct {
// 	Id        sql.NullInt64
// 	Username  sql.NullString
// 	Password  sql.NullString
// 	CreatedAt sql.NullTime
// }
// type AllUser struct {
// 	User    User
// 	Admin   Admin
// 	Mentor  Mentor
// 	Student Student
// }

func (u User) Connect() (*sql.DB, error) {
	connStr := "user=postgres dbname=Ojek sslmode=disable password=superUser host=localhost"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, err
}
