package book

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Repository interface {
	FindAll() ([]Book, error)
	FindByUserID(userID int) ([]Book, error)
	FindByID(ID int) (Book, error)
	Save(book Book) (Book, error)
	SaveBookCategory(categories []string, ID int) ([]string, error)
	Update(book Book) (Book, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Book, error) {

	rows, err := r.db.Query("SELECT * FROM books WHERE status = 'Active'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.UserID, &book.Title, &book.Writer, &book.Pages, &book.Synopsis, &book.CoverImage, &book.File, &book.Status, &book.Slug, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		// listCategory, err := r.FindCategories(book.ID)
		// if err != nil {
		// 	return nil, err
		// }
		// book.Category = listCategory
		books = append(books, book)
	}
	return books, nil
}

func (r *repository) FindByUserID(userID int) ([]Book, error) {
	rows, err := r.db.Query("SELECT * FROM books WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.UserID, &book.Title, &book.Writer, &book.Pages, &book.Synopsis, &book.CoverImage, &book.File, &book.Status, &book.Slug, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		// listCategory, err := r.FindCategories(book.ID)
		// if err != nil {
		// 	return nil, err
		// }
		// book.Category = listCategory
		books = append(books, book)
	}
	return books, nil
}

// Load relation data Category
func (r *repository) FindCategories(bookID int) (string, error) {
	row := r.db.QueryRow("SELECT c.name as category FROM book_categories as a LEFT JOIN categories as c ON a.category_id = c.id WHERE a.book_id = ?", bookID)

	var category string
		err := row.Scan(&category)
		if err != nil {
			return category, err
		}

	return category, nil

}


func (r *repository) FindByID(ID int) (Book, error){
	var book Book
	
	row := r.db.QueryRow("SELECT * FROM books WHERE id = ? ", ID)

	err := row.Scan(&book.ID, &book.UserID, &book.Title, &book.Writer, &book.Pages, &book.Synopsis, &book.CoverImage, &book.File, &book.Status, &book.Slug, &book.CreatedAt, &book.UpdatedAt)
	log.Println(err)
	if err != nil {
		return book, err
	}
	
	return book, nil
}

func (r *repository) Save(book Book) (Book, error){
	send, err := r.db.Exec("INSERT INTO books (user_id, title, writer, pages, synopsis, cover_image, file, status, slug, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?,?)", book.UserID, book.Title, book.Writer, book.Pages, book.Synopsis, book.CoverImage, book.File, book.Status, book.Slug, time.Now(), time.Now())

	if err != nil {
		return book, err
	}

	id, err := send.LastInsertId()
	if err != nil {
		return book, err
	}

	book.ID = int(id)
	
	return book, nil
}

func (r *repository) SaveBookCategory(categories []string, ID int) ([]string, error){
	valueStrings := make([]string, 0, len(categories))
    valueArgs := make([]interface{}, 0, len(categories) * 2)
    i := 0
    for _, post := range categories {
       valueStrings = append(valueStrings, "(?, ?)")
        valueArgs = append(valueArgs, ID)
		category, _ := strconv.Atoi(post)
        valueArgs = append(valueArgs, category)
        i++
    }

	 stmt := fmt.Sprintf("INSERT INTO book_categories (book_id, category_id) VALUES %s", 
                        strings.Join(valueStrings, ","))
    _, err := r.db.Exec(stmt, valueArgs...)

    return categories, err
}

func (r *repository) Update(book Book) (Book, error){
	// err := r.db.Save(&user).Error
	_, err := r.db.Exec("update books set title = ?, writer = ?, pages = ?, synopsis = ?, cover_image = ?, file = ?, status= ?, slug = ?, updated_at = ? where id = ?", book.Title, book.Writer, book.Pages, book.Synopsis, book.CoverImage, book.File, book.Status, book.Slug, time.Now(), book.ID)


	if err != nil {
		return book, err
	}

	return book, nil
}