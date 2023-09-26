package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-restfulapi-exercise/app"
	"golang-restfulapi-exercise/controller"
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/middleware"
	"golang-restfulapi-exercise/repository"
	"golang-restfulapi-exercise/service"
	"io"
	"net/http"
	"net/http/httptest"
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
	assert.Equal(t, 200, response.StatusCode)

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

}

func TestUpdateCategorySuccess(t *testing.T) {

}

func TestUpdateCategoryFailed(t *testing.T) {

}

func TestGetCategorySuccess(t *testing.T) {

}

func TestGetCategoryFailed(t *testing.T) {

}

func TestGetListCategorySuccess(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}

func TestDeleteCategoryFailed(t *testing.T) {

}
