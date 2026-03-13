package main

import (
	"ZadanieTest/Sohranenie"
	"fmt"
	"log"
)

func VivodVsexAvtorov() {
	authors, err := Sohranenie.VivodAuthors()
	if err != nil {
		log.Fatalf("Не удалось получить авторов: %v", err)
	}
	// Проверка на пустой список
	if len(authors) == 0 {
		fmt.Println("Список авторов пуст.")
		return
	}
	fmt.Println("Список авторов:")
	for _, author := range authors {
		fmt.Printf("--ID: %d | Имя: %s\n", author.ID, author.Name) //%s значит что значения в формате строки а %d Это целые числа
	}
}

func VivodKnig() {
	fmt.Print("Введите id автора для просмотра книг: ")
	var id int
	fmt.Scanln(&id)
	knigki, err := Sohranenie.VivodBook(id)
	if err != nil {
		fmt.Println("Не удалось найти автора")
	}
	if len(knigki) == 0 {
		fmt.Println("Список книг пуст.")
		return
	}
	fmt.Println("Список книг:")
	for _, knigi := range knigki {
		fmt.Printf("ID: %d | Название: %s | О чем: %s | ID Автора: %d\n", knigi.ID, knigi.Nazvaie, knigi.About, knigi.AuthorID)
	}
}

func CreateDanniiAuthor() {
	name := sborDannii("Введите Имя Автора")
	NewAuthorDanni, err := Sohranenie.NewAuthor(name)
	if err != nil {
		fmt.Print("Нет данных автора")
	}
	err = Sohranenie.DobavAuthor(NewAuthorDanni)
	if err != nil {
		fmt.Print("Не удалось сохранить данные")
		return
	}

}

func CreateDanniiBooks() {
	nazvanie := sborDannii("Введите название книги")
	about := sborDannii("Введите краткое описание книги")
	authorID := sborID("Введите id Автора")
	NewBookDanni, err := Sohranenie.NewBook(nazvanie, about, authorID)
	if err != nil {
		fmt.Print("Не достаточно данных для создания записи")
		return
	}
	err = Sohranenie.DobavBook(NewBookDanni)
	if err != nil {
		fmt.Print("Не удалось сохранить данные")
		return
	}
}

func YdalenieAuthor() {
	fmt.Print("Введите id Автора для удаления")
	var id int
	fmt.Scanln(&id)
	err := Sohranenie.DeleteAuthor(id)
	if err != nil {
		fmt.Println("Не удалось удалить Автора")
	}
}

func YdalenieBook() {
	fmt.Print("Введите id книги для удаления")
	var id int
	fmt.Scanln(&id)
	err := Sohranenie.DeliteBook(id)
	if err != nil {
		fmt.Println("Не удалось удалить книгу")
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
