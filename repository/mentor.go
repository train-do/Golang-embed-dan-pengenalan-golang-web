package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/train-do/Golang-embed-dan-pengenalan-golang-web/model"
)

// type RepoMentor struct {
// 	Admin model.Mentor
// }

//	func (r *RepoMentor) Create(tx *sql.Tx, admin *model.Admin) (int, error) {
//		query := `INSERT INTO "Admin" (user_id, name, createdat) VALUES ($1, $2, $3);`
//		err := tx.QueryRow(query, admin.User_id, admin.Name, time.Now()).Scan(&admin.Id)
//		if err != nil {
//			fmt.Println("Create Admin :", err)
//			tx.Rollback()
//			return -1, err
//		}
//		return admin.Id, nil
//	}
func CreateMentor(tx *sql.Tx, mentor *model.Mentor) (int, error) {
	query := `INSERT INTO "Mentor" (user_id, name, experience, degree, phone_number, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	err := tx.QueryRow(query, mentor.UserId, mentor.Name, mentor.Experience, mentor.Degree, mentor.PhoneNumber, time.Now()).Scan(&mentor.Id)
	if err != nil {
		fmt.Println("Create Mentor :", err)
		tx.Rollback()
		return -1, err
	}
	return mentor.Id, nil
}

func UpdateMentor(tx *sql.Tx, mentor *model.Mentor) (model.Mentor, error) {
	query := `update "Mentor" set "name"=$1, "experience"=$2, "degree"=$3, "phone_number"=$4 where "id"=$5 RETURNING *;`
	err := tx.QueryRow(query, mentor.Name, mentor.Experience, mentor.Degree, mentor.PhoneNumber, mentor.Id).Scan(&mentor.Id, &mentor.UserId, &mentor.Name, &mentor.Experience, &mentor.Degree, &mentor.PhoneNumber, &mentor.CreatedAt)
	if err != nil {
		fmt.Println("Update Mentor Error :", err)
		tx.Rollback()
		return model.Mentor{}, err
	}
	return *mentor, nil
}
