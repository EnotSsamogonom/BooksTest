package main

import (
	"ZadanieTest/Sohranenie"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:Password@localhost:1313/BooksDB")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Невозможно подключиться к базе данных: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

MenuViviod:
	for {
		vibor := Vybrano()
		switch vibor {
		case 1:
			Sohranenie.VivodAuthors()
		case 2:
			VivodKnig()
		case 3:
			CreateDanniiAuthor()
		case 4:
			CreateDanniiBooks()
		case 5:
			YdalenieAuthor()
		case 6:
			YdalenieBook()
		case 7:
			break MenuViviod
		default:
			fmt.Println("Неверно выбрана комнада.Повторите попытку")
		}
	}

}

func Vybrano() int {
	var vybor int
	fmt.Println(`Выберите необходимое действие: 
1 - Вывести всех авторов
2 - Вывести все книги по id автора
3 - Добавить Автора
4 - Добавить книгу
5 - Удалить Автора(все книги тоже удалятся)
6 - Удалить Книгу
7 - Выход из программы`)
	fmt.Scanln(&vybor)
	return vybor
}
