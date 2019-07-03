package dao

// Author @felixgravila

import (
	"database/sql"
	"github.com/felixgravila/psqlrestapi/model"
)

// GetBooks gets all books
func GetBooks(db *sql.DB) (*[]*model.Book, error) {
	var books []*model.Book
	rows, err := db.Query(`select id, isbn, title, author_id from books`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var isbn string
		var title string
		var authorid int
		err = rows.Scan(&id, &isbn, &title, &authorid)
		if err != nil {
			return nil, err
		}
		books = append(books, &model.Book{
			ID:     id,
			ISBN:   isbn,
			Title:  title,
			Author: authorid,
		})
	}

	return &books, nil
}

// GetBook gets an author by id
func GetBook(db *sql.DB, id int) (*model.Book, error) {
	rows, err := db.Query(`
           select isbn, title, author_id from books
           where id = $1`,
		id)
	if err != nil {
		return nil, err
	}
	rows.Next()
	var isbn string
	var title string
	var authorid int
	err = rows.Scan(&isbn, &title, &authorid)
	if err != nil {
		return nil, err
	}
	return &model.Book{
		ID:     id,
		ISBN:   isbn,
		Title:  title,
		Author: authorid,
	}, nil
}

// AddBook adds an author
func AddBook(db *sql.DB, isbn string, title string, author *model.Author) (*model.Book, error) {
	var id int
	err := db.QueryRow(`insert into
                books (id, isbn, title, author_id)
                values ($1 , $2, $3, $4)
                returning id`,
		id, isbn, title, author.ID).Scan(&id)
	return &model.Book{
		ID:     id,
		ISBN:   isbn,
		Title:  title,
		Author: author.ID,
	}, err
}

// DeleteBook deletes an author
func DeleteBook(db *sql.DB, id int) error {
	_, err := db.Exec(`delete from books where id = $1`, id)
	return err
}

// DeleteAllBooksByAuthor deletes all books by author
func DeleteAllBooksByAuthor(db *sql.DB, author int) error {
	_, err := db.Exec(`delete from books where author_id = $1`, author)
	return err
}

// UpdateBook updates an author
func UpdateBook(db *sql.DB, book *model.Book) error {
	_, err := db.Exec(`update books
                  set isbn = $1, title = $2, author_id = $3
                  where id = $4`,
		book.ISBN,
		book.Title,
		book.Author,
		book.ID,
	)
	return err
}
