package book

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetBooks() ([]Book, error)
	GetUserBooks(input GetBookDetailInput) ([]Book, error)
	GetBookByID(input GetBookDetailInput) (Book, error)
	CreateBook(input CreateBookInput) (Book, error)
	CreateBookCategory(input CreateBookInput, bookID int) ([]string, error)
	SaveImageCover(ID int, fileLocation string) (Book, error)
	SaveBookfile(ID int, fileLocation string) (Book, error)
	UpdateStatus(status GetBookStatusInput) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBooks() ([]Book, error) {
	books, err := s.repository.FindAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetUserBooks(input GetBookDetailInput) ([]Book, error) {

		books, err := s.repository.FindByUserID(input.ID)
		if err != nil {
			return books, err
		}

		return books, nil

}

func (s *service) GetBookByID(input GetBookDetailInput) (Book, error) {
	book, err := s.repository.FindByID(input.ID)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) CreateBook(input CreateBookInput) (Book, error) {
	book := Book{}
	book.Title = input.Title
	book.Writer = input.Writer
	book.Pages = input.Pages
	book.Synopsis = input.Synopsis
	book.Status = "Pending"
	book.UserID = input.User.ID
	
	slugString := fmt.Sprintf("%s %d", input.Title, input.User.ID)
	book.Slug = slug.Make(slugString)

	newBook, err := s.repository.Save(book)
	if err != nil {
		return newBook, err
	}
	return newBook, nil
}

func (s *service) CreateBookCategory(input CreateBookInput, bookID int) ([]string, error){

	_, err := s.repository.SaveBookCategory(input.Category, bookID)
	if err != nil {
		return nil, err
	}

	return input.Category, err
}

func (s *service) SaveImageCover(ID int, fileLocation string) (Book, error){
		// find book by id
		book, err := s.repository.FindByID(ID)
		if err != nil {
			return book, err
		}

		// update atribute ImageCover 
		book.CoverImage = fileLocation

		// save field ImageCover
		updateBook, err := s.repository.Update(book)
		if err != nil {
			return updateBook, err
		}

		return updateBook, nil
}

func (s *service) SaveBookfile(ID int, fileLocation string) (Book, error){
		// find book by id
		book, err := s.repository.FindByID(ID)
		if err != nil {
			fmt.Println("error search data book")
			return book, err
		}

		// update atribute ImageCover 
		book.File = fileLocation

		// save field ImageCover
		updateBook, err := s.repository.Update(book)
		if err != nil {
			fmt.Println("error save data book")
			return updateBook, err
		}

		return updateBook, nil
}

func (s *service) UpdateStatus(status GetBookStatusInput) (Book, error){
		// find book by id
		book, err := s.repository.FindByID(status.ID)
		if err != nil {
			fmt.Println("error search data book")
			return book, err
		}

		// update atribute Status
		book.Status = status.Status

		// save field ImageCover
		updateBook, err := s.repository.Update(book)
		if err != nil {
			fmt.Println("error save data book")
			return updateBook, err
		}

		return updateBook, nil
}