package controller

import (
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/model/web"
	"golang-restfulapi-exercise/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImplem struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImplem{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImplem) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idKategori := p.ByName("idKategori")
	id, error := strconv.Atoi(idKategori)
	helper.PanicIfError(error)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)

	// contoh web response
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplem) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())

	// contoh web response & write ke jsonnya
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplem) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// baca data dari json
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	// dari controller kita panggil servicenya
	categoryresponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)

	// contoh web response & write ke jsonnya
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
		Data:   categoryresponse,
	}

	// karena ini json maka jangan lupa agar menambahkan di header
	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImplem) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// baca data dari json

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	idKategori := p.ByName("idKategori")
	id, error := strconv.Atoi(idKategori)
	helper.PanicIfError(error)

	categoryUpdateRequest.Id = id

	// dari controller kita panggil servicenya
	categoryresponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)

	// contoh web response & write ke jsonnya
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
		Data:   categoryresponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplem) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idKategori := p.ByName("idKategori")
	id, error := strconv.Atoi(idKategori)
	helper.PanicIfError(error)

	controller.CategoryService.Delete(r.Context(), id)

	// contoh web response & write ke jsonnya
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}
