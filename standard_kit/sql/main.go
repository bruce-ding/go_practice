package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "development.sqlite3")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("open sqlite failed: %v", err)
	}
	fmt.Println("opended")

	// sqlTable := `
	// CREATE TABLE IF NOT EXISTS "users" (
	//    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
	//    "name" VARCHAR(64) NULL
	// ); `
	// _, err = db.Exec(sqlTable)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	res, err := GetAllUsers(db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(res))

	err = ExecSqlWithTransaction(db, func(tx *sql.Tx) error {
		_, err1 := tx.Exec("")
		return err1
	})
	if err != nil {
		fmt.Println(err)
	}
}

type User struct {
	ID   int
	Name sql.NullString
}

func GetAllUsers(db *sql.DB) ([]*User, error) {
	res := []*User{}
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := new(User)
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		fmt.Println(u.ID, u.Name.String)
		res = append(res, u)
	}
	return res, rows.Err()
}

func ExecSqlWithTransaction(db *sql.DB, handle func(tx *sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	if err = handle(tx); err != nil {
		return err
	}
	return tx.Commit()
}
