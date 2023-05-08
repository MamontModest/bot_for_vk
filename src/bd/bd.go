package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


var dbInfo="postgresql://postgres:password@db/"
func CreateTable() error {

    //Подключаемся к БД
    db, err := sql.Open("postgres", dbInfo)
    if err != nil {
        return err
    }
    defer db.Close()

    //Создаем таблицу users
    if _, err = db.Query(`CREATE TABLE users(ID SERIAL PRIMARY KEY, SERVICE VARCHAR(50), PASSWORD VARCHAR(50));`); err != nil {
        return err
    }
    fmt.Println("created !!")
    return nil
}