package dao

// Author @felixgravila

import (
	"database/sql"
	"github.com/felixgravila/psqlrestapi/model"
)

// GetAuthors gets all authors
func GetAuthors(db *sql.DB) (*[]*model.Author, error) {
	var authors []*model.Author
	rows, err := db.Query(`select id, firstname, lastname from authors`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var firstname string
		var lastname string
		err = rows.Scan(&id, &firstname, &lastname)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &model.Author{
			ID:        id,
			FirstName: firstname,
			LastName:  lastname,
		})
	}

	return &authors, nil
}

// GetAuthor gets an author by id
func GetAuthor(db *sql.DB, id int) (*model.Author, error) {
	rows, err := db.Query(`
           select firstname, lastname from authors
           where id = $1`,
		id)
	if err != nil {
		return nil, err
	}
	rows.Next()
	var firstname string
	var lastname string
	err = rows.Scan(&firstname, &lastname)
	if err != nil {
		return nil, err
	}
	return &model.Author{
		ID:        id,
		FirstName: firstname,
		LastName:  lastname,
	}, nil
}

// AddAuthor adds an author
func AddAuthor(db *sql.DB, firstName string, lastName string) (*model.Author, error) {
	var id int
	err := db.QueryRow(`insert into
                authors (firstname, lastname)
                values ($1 , $2)
                returning id`,
		firstName, lastName).Scan(&id)
	return &model.Author{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}, err
}

// DeleteAuthor deletes an author
func DeleteAuthor(db *sql.DB, id int) error {
	_, err := db.Exec(`delete from authors where id = $1`, id)
	return err
}

// UpdateAuthor updates an author
func UpdateAuthor(db *sql.DB, author *model.Author) error {
	_, err := db.Exec(`update authors
                  set firstname = $1, lastname = $2
                  where id = $3`,
		author.FirstName,
		author.LastName,
		author.ID)
	return err
}
