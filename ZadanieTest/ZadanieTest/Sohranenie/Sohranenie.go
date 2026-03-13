package Sohranenie

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func VivodAuthors() ([]Author, error) {
	//вывод всех авторов
	//SELECT * FROM authors
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	aut, err := db.Query(
		context.Background(),
		"SELECT id, name FROM authors ")
	if err != nil {
		return nil, err
	}
	defer aut.Close()
	var author []Author

	for aut.Next() { //проходиммся по строкам и собираем их в одно
		var a Author
		err := aut.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil, err
		}
		author = append(author, a)
	}

	if err = aut.Err(); err != nil {
		return nil, err
	}

	return author, nil
}

func VivodBook(id int) ([]Books, error) {
	//Необходим вывод книг по автору
	//SELECT * FROM books WHERE author_id = (полученное число)
	//подключаем бд
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())
	//подключилтсь

	kn, err := db.Query(
		context.Background(),
		"SELECT id, nazvanie, about, author_id FROM books WHERE author_id = ($1)",
		id)
	if err != nil {
		return nil, err
	}
	defer kn.Close()
	var knigi []Books
	//
	for kn.Next() { //проходиммся по строкам и собираем их в одно
		var b Books
		err := kn.Scan(&b.ID, &b.Nazvaie, &b.About, &b.AuthorID)
		if err != nil {
			return nil, err
		}
		knigi = append(knigi, b)
	}

	if err = kn.Err(); err != nil {
		return nil, err
	}
	//

	return knigi, nil
}

func DobavAuthor(aaa *Author) error {
	//Необходимо записать способ сохранения в бд
	//INSERT INTO authors (name) VALUES ('полученные данные')
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(
		context.Background(),
		"INSERT INTO authors (name) VALUES ($1)",
		aaa.Name)
	return err
}

func DobavBook(book *Books) error {
	//Необходимо записать способ сохранения в бд
	//INSERT INTO books (nazvanie, about, author_id) VALUES ('название', 'содержание', 1(id автора))
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(
		context.Background(),
		"INSERT INTO books (nazvanie, about, author_id) VALUES ($1, $2, $3) ",
		book.Nazvaie,
		book.About,
		book.AuthorID)
	return err
}

func DeleteAuthor(idAuthor int) error {
	//Необходимо записать способ удаления автора
	//DELETE FROM authors WHERE id = int(полученое число);

	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(
		context.Background(),
		"DELETE FROM authors WHERE id = ($1)",
		idAuthor)
	return err
}

func DeliteBook(idBook int) error {
	//Необходимо записать способ удаления книги
	//DELETE FROM books WHERE id = int(полученое число);
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(
		context.Background(),
		"DELETE FROM books WHERE id = ($1)",
		idBook)
	return err
}

//// Тут код для подключения к базе - возвращает соединение
//conn, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
//if err != nil {
//fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//os.Exit(1)
//}
//defer conn.Close(context.Background())
//
//var name string
//var weight int64
//// Тут пишешь sql-запрос и в конце в Scan записываешь полученные значения в переменные
//err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
//if err != nil {
//fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
//os.Exit(1)
//}
//
//fmt.Println(name, weight)
