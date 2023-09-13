package main

import (
	"golang-restfulapi-exercise/app"
	"golang-restfulapi-exercise/controller"
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/repository"
	"golang-restfulapi-exercise/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	// deklarasi variabelnya,1 kita ambil dari function newcategoryrepository(), 1 agi validate ambil dari dependencies,
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	// lalu kita buat service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:idKategori", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:idKategori", categoryController.Update)
	router.DELETE("/api/categories/:idKategori", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	error := server.ListenAndServe()
	helper.PanicIfError(error)
}
