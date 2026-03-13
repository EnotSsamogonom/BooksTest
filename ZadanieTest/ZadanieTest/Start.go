package main

import (
	"fmt"
)

func main() {

MenuViviod:
	for {
		vibor := Vybrano()
		switch vibor {
		case 1:
			VivodVsexAvtorov()
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
