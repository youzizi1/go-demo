package main

import "fmt"

type Book struct {
	title string
	author string
	price int	
}

type BookStore struct {
	BookList []*Book
}

func main() {
	book1 := Book{title: "Go语言学习指南1", author: "ugu1", price: 10}
	book2 := Book{title: "Go语言学习指南2", author: "ugu2", price: 20}
	book3 := Book{title: "Go语言学习指南3", author: "ugu3", price: 30}

	bookStore := BookStore{}
	bookStore.BookList = append(bookStore.BookList, &book1, &book2)
	bookStore.BookList = append(bookStore.BookList, &book3)


	for i, v := range bookStore.BookList {
		fmt.Println(i, v)
	}
}