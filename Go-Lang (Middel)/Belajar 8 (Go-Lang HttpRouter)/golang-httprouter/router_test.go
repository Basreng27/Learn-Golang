package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// ================================================== Test Router ==================================================
// func TestRouter(t *testing.T) {
// 	router := httprouter.New() //membuat router

// 	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprint(w, "Hello World")
// 	})

// 	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// 	rec := httptest.NewRecorder()

// 	router.ServeHTTP(rec, req) //ini penting
// 	res := rec.Result()

// 	bytes, _ := io.ReadAll(res.Body)
// 	assert.Equal(t, "Hello World", string(bytes)) //langsung di unitest memakai assert //harus sama hello wrodl sengn atas
// }

// ================================================== Params ==================================================
// func TestParam(t *testing.T) {
// 	router := httprouter.New() //membuat router

// 	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) { ///product/:id = membawa product yang idnya 1
// 		text := "Product " + p.ByName("id") //mengambil isi dari id
// 		fmt.Fprint(w, text)
// 	})

// 	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/product/1", nil)
// 	rec := httptest.NewRecorder()

// 	router.ServeHTTP(rec, req) //ini penting
// 	res := rec.Result()

// 	bytes, _ := io.ReadAll(res.Body)
// 	assert.Equal(t, "Product 1", string(bytes)) //langsung di unitest memakai assert //harus sama dgn atas
// }

// ================================================== Router Patterns ==================================================
// Name Parameter
// func TestNameParams(t *testing.T) {
// 	router := httprouter.New() //membuat router

// 	router.GET("/product/:id/items/:iditem", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) { ///product/:id = membawa product yang idnya 1
// 		text := "Product " + p.ByName("id") + " Item " + p.ByName("iditem") //mengambil isi dari id
// 		fmt.Fprint(w, text)
// 	})

// 	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/product/1/items/4", nil)
// 	rec := httptest.NewRecorder()

// 	router.ServeHTTP(rec, req) //ini penting
// 	res := rec.Result()

// 	bytes, _ := io.ReadAll(res.Body)
// 	assert.Equal(t, "Product 1 Item 4", string(bytes)) //langsung di unitest memakai assert //harus sama dgn atas
// }

// Catch All Parameter
func TestCacth(t *testing.T) {
	router := httprouter.New() //membuat router

	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) { ///product/:id = membawa product yang idnya 1
		text := "Image " + p.ByName("image") //mengambil isi dari cacth
		fmt.Fprint(w, text)
	})

	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/images/small/profile.png", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req) //ini penting
	res := rec.Result()

	bytes, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Image /small/profile.png", string(bytes)) //langsung di unitest memakai assert //harus sama dgn atas
}

// ================================================== Serve File ==================================================
//
//go:embed resource
var resource embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resource, "resource")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/files/hello.txt", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req) //ini penting
}

// ================================================== Panic Handler ==================================================
func TestPanic(t *testing.T) {

	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) { //membuat panic handler
		fmt.Fprint(writer, "Panic : ", error)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(body))
}

// ================================================== Not Found Handler ==================================================
// jika halaman web tdiak ketemu
func TestNotFound(t *testing.T) {

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) { //hanya mengganti menjadi handler func
		fmt.Fprint(writer, "Gak Ketemu")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/404", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Ketemu", string(body))
}

// ================================================== Method Not Allowed Handler ==================================================
func TestMethodNotAllowed(t *testing.T) {

	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) { //menambahkan methodnotallowed
		fmt.Fprint(writer, "Gak Boleh")
	})
	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "POST")
	})

	request := httptest.NewRequest("PUT", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Boleh", string(body))
}

// ================================================== Middleware ==================================================
type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive Request")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Middleware")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))

}
