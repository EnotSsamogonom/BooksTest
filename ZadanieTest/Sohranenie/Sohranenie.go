package Sohranenie

import (
	"context"
)

//DATABASE_URL=postgress://postgres:Password@localhost:1313/BooksDB

func VivodAuthors() {
	//вывод всех авторов
	//SELECT * FROM authors
}

func VivodBook(int) (_, error) {
	//Необходим вывод книг по автору
	//SELECT * FROM books WHERE author_id = (полученное число)

	if err != nil {
		return nil, err
	}
	return _, nil
}

func (book *Books) DobavBook() (_, error) {

	//Необходимо записать способ сохранения в бд
	//INSERT INTO books (nazvanie, about, author_id) VALUES ('название', 'содержание', 1(id автора))
	if err != nil {
		return nil, err
	}
	return _, nil
}

func DeliteBook(idBook int) (_, error) {
	//Необходимо записать способ удаления книги
	//DELETE FROM books WHERE id = int(полученое число);
	if err != nil {
		return nil, err
	}
	return _, nil
}

//блок
func (author *Author) DobavAuthor() (_, error) {
	//Необходимо записать способ сохранения в бд
	//INSERT INTO authors (name) VALUES ('полученные данные')
	err = db.QueryRow(
		context.Background(),
		"INSERT INTO authors (name) VALUES ($1)",
		&author.Name)
	if err != nil {
		return nil, err
	}
	return _, nil
}

//блок
func DeliteAuthor(idAuthor int) (_, error) {
	//Необходимо записать способ удаления автора
	//DELETE FROM authors WHERE id = int(полученое число);
	if err != nil {
		return nil, err
	}
	return _, nil
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
