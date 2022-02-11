package main

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Title string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s", b.Title)
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	book := Book{Title: "Game of Thrones"}
	WriteLog(book)

	count := Count(3)
	WriteLog(count)
}
