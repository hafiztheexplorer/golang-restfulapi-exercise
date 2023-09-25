package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-restfulapi-exercise/app"
	"golang-restfulapi-exercise/controller"
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/middleware"
	"golang-restfulapi-exercise/model/domain"
	"golang-restfulapi-exercise/repository"
	"golang-restfulapi-exercise/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/go-playground/assert.v1"
)

func setupTestDB() *sql.DB {
	// masukkan drivernya "mysql"
	// rumusnya seperti ini Open(driverName string, dataSourceName string)
	// datasource kalau bingung coba buka web github dari sql.DB nanti
	// akan dikasih tahu cara menulisnya
	db, error := sql.Open("mysql", "root:root2adminthistimearound@tcp(localhost:3306)/gorestful_api_exercise_test")
	helper.PanicIfError(error)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func setupRouter(db *sql.DB) http.Handler {

	// deklarasi variabelnya,1 kita ambil dari function newcategoryrepository(), 1 agi validate ambil dari dependencies,
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	// lalu kita buat service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

// skenario test

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	// setup router
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"namakategori":"test3"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY2") // pakai apikey random lain yg tidak dispesifikkan

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
	// assert.Equal(t, "test3", responseBody["data"].(map[string]interface{})["namakategori"])
}

func TestCreateCategorySuccess(t *testing.T) {

	db := setupTestDB()
	truncateCategory(db)
	// setup router
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"namakategori":"test3"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "test3", responseBody["data"].(map[string]interface{})["namakategori"])

}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	// setup router
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"namakategori":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "INVALID/BAD REQUEST", responseBody["status"])

}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	// untuk update harus create data dulu menggunakan repository
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Namakategori: "input from unit test 1",
	})
	tx.Commit()

	// setup router
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"namakategori":"input from unit test 1"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "input from unit test 1", responseBody["data"].(map[string]interface{})["namakategori"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	// untuk update harus create data dulu menggunakan repository
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Namakategori: "input from unit test 1",
	})
	tx.Commit()

	// setup router
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"namakategori":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "INVALID/BAD REQUEST", responseBody["status"])
	// assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// assert.Equal(t, "input from unit test 1", responseBody["data"].(map[string]interface{})["namakategori"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	// untuk update harus create data dulu menggunakan repository
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Namakategori: "input from unit test 1",
	})
	tx.Commit()

	// setup router
	router := setupRouter(db)
	// requestBody := strings.NewReader(`{"namakategori":"input from unit test 1"}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	// request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Namakategori, responseBody["data"].(map[string]interface{})["namakategori"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	// untuk update harus create data dulu menggunakan repository
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Namakategori: "input from unit test 1",
	})
	tx.Commit()

	// setup router
	router := setupRouter(db)
	// requestBody := strings.NewReader(`{"namakategori":"input from unit test 1"}`)
	// perhatikan di sini sengaja kita get category di id yang ngga ada dari logic diatas
	// itu mengenerate id 0, nah sedangkan yang di get id existing+1
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id+1), nil)
	// request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "DATA NOT FOUND", responseBody["status"])

}

func TestGetListCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	// untuk update harus create data dulu menggunakan repository
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()

	category1 := categoryRepository.Create(context.Background(), tx, domain.Category{Namakategori: "input from unit test ke-1"})
	category2 := categoryRepository.Create(context.Background(), tx, domain.Category{Namakategori: "input from unit test ke-2"})

	tx.Commit()

	// setup router
	router := setupRouter(db)
	// requestBody := strings.NewReader(`{"namakategori":"input from unit test 1"}`)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	// request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var categories = responseBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Namakategori, categoryResponse1["namakategori"])

	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Namakategori, categoryResponse2["namakategori"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	// untuk update harus create data dulu menggunakan repository
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Namakategori: "input from unit test 1",
	})
	tx.Commit()

	// setup router
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"namakategori":"input from unit test 1"}`)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	// assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// assert.Equal(t, "input from unit test 1", responseBody["data"].(map[string]interface{})["namakategori"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	// untuk update harus create data dulu menggunakan repository
	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, domain.Category{
		Namakategori: "input from unit test 2",
	})
	tx.Commit()

	// setup router
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"namakategori":"input from unit test 2"}`)
	// di bawh ini kita coba delete category id yang tidak pernah ada, category.Id + 1
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id+1), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "APIKEY")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode) // cek apakah response 200

	// untuk check kondisi return http code lain
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// untuk cek response body nya apa aja
	fmt.Println(responseBody)

	// lalu kita assert
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "DATA NOT FOUND", responseBody["status"])
	// assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	// assert.Equal(t, "input from unit test 1", responseBody["data"].(map[string]interface{})["namakategori"])
}
