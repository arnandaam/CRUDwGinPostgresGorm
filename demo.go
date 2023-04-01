package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// database -> db_go_sql
// schema -> grouping -> schema user, schema product, schema order
// table -> table di schema user: user_profile, user_documents
// index, function, etc

// create database db_go_sql;
// create schema db_go_sql;
// create table db_go_sql.employees (
// 	id serial primary key,
// 	full_name varchar(50) not null,
// 	email text unique not null,
// 	age int not null,
// 	division varchar(20) not null
// );

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "4rn4nd44m"
// 	dbname   = "db_book"
// )

// var (
// 	db  *sql.DB
// 	err error
// )

// type Book struct {
// 	BookID string
// 	Title  string
// 	Author string
// 	Decs   string
// }

// func main() {
// 	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err = sql.Open("postgres", dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected to database")

// 	//CreateEmployee()
// 	GetEmployees()
// 	//UpdateEmployee()
// }

// // query INSERT, UPDATE, DELETE -> Exec
// // query SELECT -> Query/QueryRow

// func CreateEmployee() {
// 	var Book = Book{}

// 	sqlStatement := `
// 	INSERT INTO books
// 	(
// 		title,
// 		author,
// 		des
// 	)
// 	VALUES (
// 		$1, $2, $3
// 	)
// 	RETURNING *
// 	`

// 	err = db.QueryRow(sqlStatement, "Cara Sukses 2k23", "cahyo", "mengapai mimpi dengan sukses").
// 		Scan(&Book.BookID, &Book.Title, &Book.Author, &Book.Decs)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("buku di tambah", Book)
// }

// func GetEmployees() {
// 	var result = []Book{}
// 	query := `SELECT * FROM books ORDER BY id`

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var temp = Book{}
// 		err = rows.Scan(&temp.BookID, &temp.Title, &temp.Author, &temp.Decs)
// 		if err != nil {
// 			panic(err)
// 		}

// 		result = append(result, temp)
// 	}

// 	fmt.Println("Result", result)
// }

// func UpdateEmployee() {
// 	query := `
// 	UPDATE db_go_sql.employees
// 	SET
// 		title = $2,
// 		author = $3,
// 		des= $4
// 	WHERE
// 		id = $1;
// 	`

// 	res, err := db.Exec(query, 1, "Anang", "anang@mail.com", "Marketing")
// 	if err != nil {
// 		panic(err)
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Rows affected", count)
// }
