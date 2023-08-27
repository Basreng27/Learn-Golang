package main

import (
	"golangapi/app"
	"golangapi/controller"
	"golangapi/helper"
	"golangapi/repository"
	"golangapi/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()                                              //membuat router
	router.GET("/api/categories/", categoryController.FindAll)              //sesuai file di json
	router.GET("/api/categories/:categoryId", categoryController.FindById)  //sesuai file di json
	router.POST("/api/categories/", categoryController.Create)              //sesuai file di json
	router.PUT("/api/categories/:categoryId", categoryController.Update)    //sesuai file di json
	router.DELETE("/api/categories/:categoryId", categoryController.Delete) //sesuai file di json

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
