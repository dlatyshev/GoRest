package models

import "slices"

var db []Book

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	db = []Book{
		{Id: 1, Title: "2001: A Space Odyssey", Author: Author{Id: 1, Name: "Arthur C.", LastName: "Clarke", BornYear: 1917}, YearPublished: 2020},
		{Id: 2, Title: "Foundation", Author: Author{Id: 2, Name: "Isaac", LastName: "Asimov", BornYear: 1920}, YearPublished: 2021},
		{Id: 3, Title: "The Time Machine", Author: Author{Id: 3, Name: "H.G.", LastName: "Wells", BornYear: 1866}, YearPublished: 2022},
		{Id: 4, Title: "The Hobbit", Author: Author{Id: 4, Name: "J.R.R.", LastName: "Tolkien", BornYear: 1892}, YearPublished: 2023},
		{Id: 5, Title: "The Canterbury Tales", Author: Author{Id: 5, Name: "Geoffrey", LastName: "Chaucer", BornYear: 1343}, YearPublished: 2024},
	}
}

func FindBookById(id int) (Book, bool) {
	for _, book := range db {
		if book.Id == id {
			return book, true
		}
	}
	return Book{}, false
}

func GetAllBooks() []Book {
	return db
}

func AddBook(book Book) {
	db = append(db, book)
}

func UpdateBook(id int, updatedBook Book) bool {
	for i := range db {
		if db[i].Id == id {
			// Update the book details directly in the slice
			db[i].Title = updatedBook.Title
			db[i].YearPublished = updatedBook.YearPublished
			// Update Author fields individually, preserving the ID
			db[i].Author.Name = updatedBook.Author.Name
			db[i].Author.LastName = updatedBook.Author.LastName
			db[i].Author.BornYear = updatedBook.Author.BornYear
			return true // Book found and updated
		}
	}
	return false // Book not found
}

func DeleteBook(id int) {
	for i, b := range db {
		if b.Id == id {
			db = slices.Delete(db, i, i+1)
			return
		}
	}
}
