package middleware

import (
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

// buat function untuk membuat handler di atas
func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {

	return &AuthMiddleware{Handler: handler}

}

// membuat func sesuai dengan kontrak handler yang ada di server
// apabila ada request masuk akan diproses di sini dulu,
func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// apakah apikey sudah benar
	if r.Header.Get("X-API-Key") == "APIKEY" {
		//ok
		middleware.Handler.ServeHTTP(w, r)
	} else {
		//error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.Webresponse{
			Code:   http.StatusUnauthorized, // internal server error codenya 500
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
