package main

import (
	"ZadanieTest/Sohranenie"
	"fmt"
)

func CreateDanniiBooks() {
	nazvanie := sborDannii("Введите название книги")
	about := sborDannii("Введите краткое описание книги")
	authorID := sborID("Введите id Автора")
	NewBookDanni, err := Sohranenie.NewBook(nazvanie, about, authorID)
	if err != nil {
		fmt.Print("Не достаточно данных для создания записи")
		return
	}
	_, err := Sohranenie.DobavBook(NewBookDanni)
	if err != nil {
		fmt.Print("Не удалось сохранить данные")
		return
	}
}

func CreateDanniiAuthor() {
	name := sborDannii("Введите Имя Автора")
	NewAuthorDanni, err := Sohranenie.NewAuthor(name)
	if err != nil {
		fmt.Print("Нет данных автора")
	}
	_, err := Sohranenie.DobavAuthor(NewAuthorDanni)
	if err != nil {
		fmt.Print("Не удалось сохранить данные")
		return
	}

}

func VivodKnig() {
	fmt.Print("Введите id автора для прросмотра книг")
	var id int
	fmt.Scanln(id)
	_, err := Sohranenie.VivodBook(id)
	if err != nil {
		fmt.Println("Не удалось найти автора")
	}
}

func YdalenieBook() {
	fmt.Print("Введите id книги для удаления")
	var id int
	fmt.Scanln(id)
	_, err := Sohranenie.DeliteBook(id)
	if err != nil {
		fmt.Println("Не удалось удалить книгу")
	}

}

func YdalenieAuthor() {
	fmt.Print("Введите id Автора для удаления")
	var id int
	fmt.Scanln(id)
	_, err := Sohranenie.DeliteAuthor(id)
	if err != nil {
		fmt.Println("Не удалось удалить Автора")
	}
}

func sborDannii(aaa string) string {
	fmt.Print(aaa + ": ")
	var result string
	fmt.Scanln(&result)
	return result
}

func sborID(sss string) int {
	fmt.Print(sss + ": ")
	var result int
	fmt.Scanln(&result)
	return result
}
