package controller

import (
	"encoding/json"
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

	categoryResponse := controller.CategoryService.FindByIdGet(r.Context(), id)

	// contoh web response
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	encoder := json.NewEncoder(w)
	error = encoder.Encode(webResponse)
	helper.PanicIfError(error)

	// karena ini json maka jangan lupa agar menambahkan di header
	w.Header().Add("Content-Type", "application/json")
}

func (controller *CategoryControllerImplem) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAllGet(r.Context())

	// contoh web response & write ke jsonnya
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	encoder := json.NewEncoder(w)
	error := encoder.Encode(webResponse)
	helper.PanicIfError(error)

	// karena ini json maka jangan lupa agar menambahkan di header
	w.Header().Add("Content-Type", "application/json")
}

func (controller *CategoryControllerImplem) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// baca data dari json
	decoder := json.NewDecoder(r.Body)

	categoryCreateRequest := web.CategoryCreateRequest{}
	error := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(error)

	// dari controller kita panggil servicenya
	categoryresponse := controller.CategoryService.CreatePost(r.Context(), categoryCreateRequest)

	// contoh web response & write ke jsonnya
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
		Data:   categoryresponse,
	}

	encoder := json.NewEncoder(w)
	error = encoder.Encode(webResponse)
	helper.PanicIfError(error)

	// karena ini json maka jangan lupa agar menambahkan di header
	w.Header().Add("Content-Type", "application/json")
}

func (controller *CategoryControllerImplem) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// baca data dari json
	decoder := json.NewDecoder(r.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	error := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(error)

	idKategori := p.ByName("idKategori")
	id, error := strconv.Atoi(idKategori)
	helper.PanicIfError(error)

	categoryUpdateRequest.Id = id

	// dari controller kita panggil servicenya
	categoryresponse := controller.CategoryService.UpdatePut(r.Context(), categoryUpdateRequest)

	// contoh web response & write ke jsonnya
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
		Data:   categoryresponse,
	}

	encoder := json.NewEncoder(w)
	error = encoder.Encode(webResponse)
	helper.PanicIfError(error)

	// karena ini json maka jangan lupa agar menambahkan di header
	w.Header().Add("Content-Type", "application/json")
}

func (controller *CategoryControllerImplem) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idKategori := p.ByName("idKategori")
	id, error := strconv.Atoi(idKategori)
	helper.PanicIfError(error)

	controller.CategoryService.DeleteDelete(r.Context(), id)

	// contoh web response & write ke jsonnya
	webResponse := web.Webresponse{
		Code:   200,
		Status: "OK",
	}

	encoder := json.NewEncoder(w)
	error = encoder.Encode(webResponse)
	helper.PanicIfError(error)

	// karena ini json maka jangan lupa agar menambahkan di header
	w.Header().Add("Content-Type", "application/json")

}
