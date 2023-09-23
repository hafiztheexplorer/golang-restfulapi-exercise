package main

import (
	"golang-restfulapi-exercise/app"
	"golang-restfulapi-exercise/controller"
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/middleware"
	"golang-restfulapi-exercise/repository"
	"golang-restfulapi-exercise/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	// deklarasi variabelnya,1 kita ambil dari function newcategoryrepository(), 1 agi validate ambil dari dependencies,
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	// lalu kita buat service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	error := server.ListenAndServe()
	helper.PanicIfError(error)
}
