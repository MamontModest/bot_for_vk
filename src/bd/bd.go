package bd

import (
	"database/sql"
	"fmt"
	"github.com/MamontModest/bot_for_vk/model"
	_ "github.com/lib/pq"
	"log"
)

var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", "localhost", "5432", "postgres", "QWertas1122", "bot_vk", "disable")

func CreateTables() error {

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()
	if _, err = db.Query(`CREATE TABLE if not exists users(ID BIGINT, SERVICE VARCHAR(50),LOGIN VARCHAR(50), PASSWORD VARCHAR(50), PRIMARY KEY(ID, SERVICE));`); err != nil {
		return err
	}
	log.Println("Created database users!!")
	return nil
}

func createConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected!")
	return db, nil
}

func CreatePassword(u *model.User) error {
	db, err := createConnection()
	if err != nil {
		return err
	}
	if _, err := db.Query(`INSERT INTO users(ID, SERVICE, LOGIN, PASSWORD) values ($1, $2, $3, $4);`, u.Uid, u.Service, u.Login, u.Password); err != nil {
		log.Println("Create service error", err, u.Uid)
		return err
	}
	return nil
}

func DeletePassword(u *model.User) error {
	db, err := createConnection()
	if err != nil {
		return err
	}
	if _, err := db.Query(`DELETE FROM users where ID=$1 and SERVICE=$2;`, u.Uid, u.Service); err != nil {
		log.Println("Delete service error", err, u.Uid)
		return err
	}
	return nil
}

func SearchPassword(u *model.User) error {
	db, err := createConnection()
	if err != nil {
		return err
	}
	if err := db.QueryRow(
		`SELECT LOGIN, PASSWORD FROM users where ID=$1 and SERVICE=$2;`,
		u.Uid, u.Service).Scan(&u.Login, &u.Password); err != nil {
		log.Println("Search service error", err, u.Uid)
		return err
	}
	return nil
}
