package handler

import (
	"diary/auth"
	"diary/helper"
	"diary/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	// aambil input dari user
	// ,ap inputan ke struct Register
	// passing struc sebagai parameter

	var input user.RegisterUserInput
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		 c.JSON(http.StatusUnprocessableEntity, response)
		 return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	// token, err := h.authService.GenerateToken(newUser.ID)
	// if err != nil {
	// 	response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	 return
	// }

	formatter := user.FormatRegisterUser(newUser)
	
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}	


func (h *userHandler) Login(c *gin.Context) {
	// user menginput data
	// inputan user diambil handler
	// mapping input user ke input struct
	// input struct kirim ke service
	// service mencari data user dengan bantuan repository
	// match password


	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		 c.JSON(http.StatusUnprocessableEntity, response)
		 return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		 c.JSON(http.StatusUnprocessableEntity, response)
		 return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Login Successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailRegister(c *gin.Context) {
	// masuk inputan email dari user
	// inputan di mapping ke struct input
	// struct input dikirim ke service
	// service akan check email sudah tersedia aatau belum lewat repository

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		 c.JSON(http.StatusUnprocessableEntity, response)
		 return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}

		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		 c.JSON(http.StatusUnprocessableEntity, response)
		 return
	}	

	data := gin.H{
		"is_email_available" : IsEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if IsEmailAvailable{
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}


func (h *userHandler) UploadAvatar(c *gin.Context) {
	// ambil inputan form dari user
	// simpan gambar di folder images
	// find id user di service  JWT
	// update data field file_avatar dengan repository

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	// id by jwt
	userID := 3

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		 return
	}

	data := gin.H{"is_uploaded": true}

	response := helper.APIResponse("Successfully upload image", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)



}