package handler

import (
	"diary/book"
	"diary/helper"
	"diary/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

func (h *bookHandler) GetBooks(c *gin.Context) {

	books, err := h.service.GetBooks()
	if err != nil {
		response := helper.APIResponse("Error to get books data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Books", http.StatusOK, "success", book.FormatBooks(books))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) GetUserBooks(c *gin.Context) {
	// userID, _ := strconv.Atoi(c.Query("user_id"))
	var input book.GetBookDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get books data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	books, err := h.service.GetUserBooks(input)
	if err != nil {
		response := helper.APIResponse("Error to get books data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Books", http.StatusOK, "success", book.FormatBooks(books))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) GetBook(c *gin.Context){
	var input book.GetBookDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get books data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	bookDetail, err := h.service.GetBookByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get books data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of Book", http.StatusOK, "success", book.FormatBookDetail(bookDetail))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) CreateBook(c *gin.Context){
	var input book.CreateBookInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create new book", http.StatusUnprocessableEntity, "error", errorMessage)
		 c.JSON(http.StatusUnprocessableEntity, response)
		 return
	}

	currentUser := c.MustGet("CurrentUser").(user.User)
	
	input.User = currentUser

	newBook, err := h.service.CreateBook(input)
	if err != nil {
		response := helper.APIResponse("Failed to create new book", http.StatusBadRequest, "error", nil)
		 c.JSON(http.StatusBadRequest, response)
		 return
	}

	newCategory, err := h.service.CreateBookCategory(input, newBook.ID)
	if err != nil {
		response := helper.APIResponse("Failed to create book's category", http.StatusBadRequest, "error", newCategory)
		 c.JSON(http.StatusBadRequest, response)
		 return
	}

	response := helper.APIResponse("Success to create new book", http.StatusOK, "success", book.FormatBook(newBook))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) SaveImageCover(c *gin.Context) {

	var input book.GetBookDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to upload image uri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("image_cover")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image form file", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	// id by jwt
	currentUser := c.MustGet("CurrentUser").(user.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/cover/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	log.Println(err)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image upload", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	_, err = h.service.SaveImageCover(input.ID, path)
	log.Println(err)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image send data", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Successfully upload image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}


func (h *bookHandler) Savefile(c *gin.Context) {

	var input book.GetBookDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to upload file", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload file", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	// id by jwt
	currentUser := c.MustGet("CurrentUser").(user.User)
	userID := currentUser.ID

	path := fmt.Sprintf("bookFiles/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload file", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	_, err = h.service.SaveBookfile(input.ID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image send data", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Successfully upload image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *bookHandler) GetReadBook(c *gin.Context){
	var input book.GetBookDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get books data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	bookData, err := h.service.GetBookByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get books data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of Book", http.StatusOK, "success", book.FormatReadBook(bookData))
	c.JSON(http.StatusOK, response)
}

func (h *bookHandler) UpdateBookStatus(c *gin.Context){
	var input book.GetBookStatusInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update book status", http.StatusUnprocessableEntity, "error", errorMessage)
		 c.JSON(http.StatusUnprocessableEntity, response)
		 return
	}

	updateBook, err := h.service.UpdateStatus(input)
	if err != nil {
		response := helper.APIResponse("Failed to update book status", http.StatusBadRequest, "error", nil)
		 c.JSON(http.StatusBadRequest, response)
		 return
	}

	response := helper.APIResponse("Success to create new book", http.StatusOK, "success", book.FormatBook(updateBook))
	c.JSON(http.StatusOK, response)
}