package app

import (
	"golang-restfulapi-exercise/controller"
	"golang-restfulapi-exercise/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:idKategori", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:idKategori", categoryController.Update)
	router.DELETE("/api/categories/:idKategori", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
