package web_server

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

//================================================== Membuat Server ==================================================
// func TestWebServer(t *testing.T) {
// 	//=============== membuat server untuk menjalankan ===============
// 	server := http.Server{
// 		Addr: "localhost:8080",
// 	}
// 	//=============== membuat server untuk menjalankan ===============

// 	err := server.ListenAndServe() //untuk menjalankan servernya
// 	if err != nil {
// 		panic(err)
// 	}
// 	//setelah di runing di termminal kemudian masuk ke web dengan ketikan "localhost:8080"
// }

// ================================================== Handler ==================================================
// func TestHandler(t *testing.T) {
// 	//membuat hwllo world di web
// 	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) { //membuat handler //writer = response yang akan diberikan ke client //request = yang dikirim oleh client
// 		// logic web
// 		fmt.Fprint(writer, "Hello World") //println = akan menampilkan ke console //jika ingin di web Fprint = ke writer
// 	}

// 	server := http.Server{ //membuat server
// 		Addr:    "localhost:8080",
// 		Handler: handler,
// 	}

// 	err := server.ListenAndServe() //memanggil
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ================================================== ServeMux ==================================================
// // untuk membuat url kurang lebih
// func TestServeMux(t *testing.T) {
// 	mux := http.NewServeMux() //membuat mux baru atau tempat penampungan handler

// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //w = writer, r = request
// 		fmt.Fprint(w, "Hello World")
// 	})

// 	mux.HandleFunc("/gaiz", func(w http.ResponseWriter, r *http.Request) { //'/gaiz' = URL /endpoint = harus berbeda
// 		fmt.Fprint(w, "Hello Gaiz")
// 	})

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: mux,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ================================================== Request ==================================================
// func TestRequest(t *testing.T) {
// 	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, r.Method) //seperti GET, POST, PUT, dll
// 		fmt.Fprint(w, r.RequestURI)
// 	}

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: handler,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ================================================== HTTP Test ==================================================
// // biasanya di pakai unit test untuk web
// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World")
// }

// func TestHelloHandler(t *testing.T) {
// 	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil) //tidak perlu menggunakan server //nil = bisa diisi body
// 	recorder := httptest.NewRecorder()

// 	HelloHandler(recorder, request)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)

// 	fmt.Println(response.StatusCode)
// 	fmt.Println(response.Status)
// 	fmt.Println(string(body))
// }

// ================================================== Query Parameter ==================================================
// // Query Parameter
// func SayHello(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("name") //r.URL = untuk mendapatkan data struct URLnya, Query() = untuk mengambil data querynya, Get("name") = mengambil data berdasarkan nama query nya
// 	if name == "" {
// 		fmt.Fprint(w, "Hello")
// 	} else {
// 		fmt.Fprintf(w, "Hello %s", name) //%s = string, mendukung formater
// 	}
// }
// func TestQueryParameter(t *testing.T) {
// 	reqest := httptest.NewRequest("GET", "https://localhost:8080/hello?name=wandi", nil) //menambahkan hello?name=wandi untuk mengambil data dengan nama wandi
// 	recorder := httptest.NewRecorder()

// 	SayHello(recorder, reqest)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)

// 	fmt.Println(string(body))
// }

// // Multiple Query Parameter
// func MultipleQuery(w http.ResponseWriter, r *http.Request) {
// 	firstname := r.URL.Query().Get("firstname")
// 	lastname := r.URL.Query().Get("lastname")

// 	fmt.Fprintf(w, "Nama Depan %s . Nama Belakang %s", firstname, lastname)
// }

// func TestMultipleQuery(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello?firstname=Rayandra&lastname=WandiMarselana", nil)
// 	record := httptest.NewRecorder()

// 	MultipleQuery(record, *request)
// 	response := record.Result()
// 	body, _ := io.ReadAll(response.Body)

// 	fmt.Println(string(body))
// }

// // Multiple Value Query Parameter
// // mendapatkan semua isi di url
// func MultipleValueQueryParameter(w http.ResponseWriter, r http.Request) {
// 	var query url.Values = r.URL.Query()      //mengambil yang pertama
// 	var names []string = query["name"]        //mengambil semua data yang isinya key name (struct)
// 	fmt.Fprintln(w, strings.Join(names, " ")) //menggabungkan datanya
// }

// func TestMultipleValueQueryParameter(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello?name=Rayandra&name=Wandi&name=Marselana", nil)
// 	record := httptest.NewRecorder()

// 	MultipleValueQueryParameter(record, *request)
// 	response := record.Result()
// 	body, _ := io.ReadAll(response.Body)

// 	fmt.Println(string(body))
// }

// ================================================== Header ==================================================
// // Request
// func RequestHeader(w http.ResponseWriter, r *http.Request) {
// 	contentType := r.Header.Get("content-type") //menangkap header
// 	fmt.Fprint(w, contentType)
// }

// func TestRequestHeader(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// 	request.Header.Add("conten-type", "application/json")
// 	recorder := httptest.NewRecorder()

// 	RequestHeader(recorder, request)

// 	response := recorder.Result() //membaca recorder
// 	body, _ := io.ReadAll(response.Body)
// 	fmt.Println(string(body))
// }

// // Response
// func ResponseHeader(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Powered", "Rayandra Wandi Marselana")
// 	fmt.Fprint(w, "OK")
// }

// func TestResponseHeader(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// 	recorder := httptest.NewRecorder()

// 	ResponseHeader(recorder, request)

// 	poweresby := recorder.Header().Get("Powered") //memanggil
// 	fmt.Println(poweresby)
// }

// ================================================== Form Post ==================================================
// // Request.PostForm
// func FormPost(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm() //parseform terlebihdahulu
// 	if err != nil {
// 		panic(err)
// 	}

// 	// r.PostFormValue("firstname")//langsung mengambil tanpa harus parseform

// 	//mengambil data
// 	firstname := r.PostForm.Get("firstname")
// 	lastname := r.PostForm.Get("lastname")
// 	fmt.Fprintf(w, "Hallo %s %s", firstname, lastname)
// }

// func TestFormPost(t *testing.T) {
// 	bodyrequest := strings.NewReader("firstname=Rayandra&lastname=Wandi") //mirip seperti query parameter tapi disimpan di body
// 	request := httptest.NewRequest(http.MethodPost, "https://localhost:8080/", bodyrequest)
// 	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	recorder := httptest.NewRecorder()

// 	FormPost(recorder, request)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)
// 	fmt.Println(string(body))
// }

// ================================================== Respose Code ==================================================
// func ResponseCode(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("name")
// 	if name == "" {
// 		w.WriteHeader(http.StatusBadRequest) //bad request //untuk response code
// 		fmt.Fprint(w, "Name is empty")
// 	} else {
// 		w.WriteHeader(200) //success request
// 		fmt.Fprintf(w, "Success Hi %s", name)
// 	}
// }

// func TestResponseCode(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// 	recorder := httptest.NewRecorder()

// 	ResponseCode(recorder, request)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)
// 	fmt.Println(response.StatusCode) //mengetahui brp status response code
// 	fmt.Println(response.Status)
// 	fmt.Println(string(body))
// }

// ================================================== Cookie ==================================================
// // set cookie
// func SetCookie(w http.ResponseWriter, r *http.Request) {
// 	cookie := new(http.Cookie)               //membuat cookie
// 	cookie.Name = "X-ZPN-Name"               //nama cookie
// 	cookie.Value = r.URL.Query().Get("name") //isi dari cookie
// 	cookie.Path = "/"                        //aktif di url mana

// 	http.SetCookie(w, cookie)
// 	fmt.Fprint(w, "Success Create Cookie")
// }

// // membaca cookie / mengambil
// func GetCookie(w http.ResponseWriter, r *http.Request) {
// 	cookie, err := r.Cookie("X-ZPN-Name")
// 	if err != nil {
// 		fmt.Fprint(w, "No Cookie") //coolie tidak ada
// 	} else {
// 		fmt.Fprintf(w, "Hi %s", cookie.Value) //mengambil isi dari cookie
// 	}
// }

// // coba cookie
// func TestCookie(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/set-cookie", SetCookie)
// 	mux.HandleFunc("/get-cookie", GetCookie)

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: mux,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// // unitest cookie
// func TestUnitestSetCookie(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/?name=wandi", nil)
// 	recorder := httptest.NewRecorder()

// 	SetCookie(recorder, request)

// 	cookies := recorder.Result().Cookies()

// 	for _, cookie := range cookies {
// 		fmt.Printf("%s : %s", cookie.Name, cookie.Value)
// 	}

// }

// func TestUnitestGetCookie(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// 	cookie := new(http.Cookie)
// 	cookie.Name = "X-PZN-Name"
// 	cookie.Value = "Wandi"
// 	request.AddCookie(cookie)

// 	recorder := httptest.NewRecorder()

// 	GetCookie(recorder, request)

// 	body, _ := io.ReadAll(recorder.Result().Body)

// 	fmt.Println(string(body))

// }

// ================================================== File Server ==================================================
// // // contoh isinya ada di folder resource
// // func TestFileServer(t *testing.T) {
// // 	directory := http.Dir("./resource")      //tempat penyimpanan file
// // 	fileserver := http.FileServer(directory) //file server

// // 	mux := http.NewServeMux()
// // 	// mux.Handle("/static/", fileserver)
// // 	mux.Handle("/static/", http.StripPrefix("/static", fileserver)) //menghapus "/static"

// // 	server := http.Server{
// // 		Addr:    "localhost:8080",
// // 		Handler: mux,
// // 	}

// // 	err := server.ListenAndServe()
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // }

// // goembed
// //
// //go:embed resource
// var resource embed.FS

// func TestFileServerGoEmbed(t *testing.T) {
// 	directory, _ := fs.Sub(resource, "resource") //jika tidak ketemu maka balikan error
// 	fileServer := http.FileServer(http.FS(directory))

// 	mux := http.NewServeMux()
// 	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: mux,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ================================================== ServeFile ==================================================
// func ServeFile(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Query().Get("name") != "" {
// 		http.ServeFile(w, r, "./resource/ok.html")
// 	} else {
// 		http.ServeFile(w, r, "./resource/notfound.html")
// 	}
// }

// func TestServeFile(t *testing.T) {
// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: http.HandlerFunc(ServeFile),
// 	}
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ================================================== Template ==================================================
// // menggunakan txt
// // func SimpleHtmlTemplate(w http.ResponseWriter, r *http.Request) {
// // 	templatetext := `<html><body>{{.}}</body></html>` //template html di golang

// // 	// t := template.Must(template.New("SIMPLE").Parse(templatetext))//sama seperti yang bawah
// // 	t, err := template.New("SIMPLE").Parse(templatetext) //template.New("SIMPLE") = nama template
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
// // }

// // func TestHtmlTemplate(t *testing.T) {
// // 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// // 	recorder := httptest.NewRecorder()

// // 	SimpleHtmlTemplate(recorder, request)

// // 	response := recorder.Result()
// // 	body, _ := io.ReadAll(response.Body)
// // 	fmt.Println(string(body))
// // }

// // menggunakan file
// // di folder template
// // func FileHtmlTemplate(w http.ResponseWriter, r *http.Request) {
// // 	t := template.Must(template.ParseFiles("./template/simple.gohtml")) //mengambil file htmlnya dengan format gohtml

// // 	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML File Template")
// // }

// // func TestFileHtmlTemplate(t *testing.T) {
// // 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// // 	recorder := httptest.NewRecorder()

// // 	FileHtmlTemplate(recorder, request)

// // 	response := recorder.Result()
// // 	body, _ := io.ReadAll(response.Body)
// // 	fmt.Println(string(body))
// // }

// // Directory template
// // func DirectoryFileHtmlTemplate(w http.ResponseWriter, r *http.Request) {
// // 	t := template.Must(template.ParseGlob("./template/*.gohtml")) //menggunakan PArseGlob

// // 	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML File Template")//jika namanya tidak ada maka tidak akan keluar isi templatenya
// // }

// // func TestDirectoryFileHtmlTemplate(t *testing.T) {
// // 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// // 	recorder := httptest.NewRecorder()

// // 	DirectoryFileHtmlTemplate(recorder, request)

// // 	response := recorder.Result()
// // 	body, _ := io.ReadAll(response.Body)
// // 	fmt.Println(string(body))
// // }

// // Embed
// //
// //go:embed template/*.gohtml
// var templates embed.FS

// func FileHtmlTemplateEmbed(w http.ResponseWriter, r *http.Request) {
// 	t := template.Must(template.ParseFS(templates, "template/*.gohtml")) //Menggunakan ParseFS

// 	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML File Template")
// }

// func TestFileHtmlTemplateEmbed(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// 	recorder := httptest.NewRecorder()

// 	FileHtmlTemplateEmbed(recorder, request)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)
// 	fmt.Println(string(body))
// }

// ================================================== Template Data ==================================================
// // Menggunakan Map
// // func TemplateData(w http.ResponseWriter, r *http.Request) {
// // 	t := template.Must(template.ParseFiles("./template/name.gohtml"))

// // 	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
// // 		"Title": "Belajar",
// // 		"Name":  "Rayandra Wandi Marselana",
// // 	})
// // }

// // func TestTemplateData(t *testing.T) {
// // 	req := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
// // 	rec := httptest.NewRecorder()

// // 	TemplateData(rec, req)

// // 	res := rec.Result()
// // 	body, _ := io.ReadAll(res.Body)

// // 	fmt.Println(string(body))
// // }

// // Menggunakan Struct
// type Addres struct {
// 	Street string
// }

// type Page struct {
// 	Title, Name string
// 	Addres      Addres
// }

// func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
// 	t := template.Must(template.ParseFiles("./template/name.gohtml"))

// 	t.ExecuteTemplate(w, "name.gohtml", Page{
// 		Title: "Belajar",
// 		Name:  "Rayandra Wandi Marselana",
// 		Addres: Addres{
// 			Street: "Cilengkrang",
// 		},
// 	})
// }

// func TestTemplateDataStruct(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
// 	rec := httptest.NewRecorder()

// 	TemplateDataStruct(rec, req)

// 	res := rec.Result()
// 	body, _ := io.ReadAll(res.Body)

// 	fmt.Println(string(body))
// }

// ================================================== Template Action ==================================================
// // seleksi kondisi
// // func TemplateActionSeleksi(w http.ResponseWriter, r *http.Request) {
// // 	t := template.Must(template.ParseFiles("./template/templateactionseleksi.gohtml"))

// // 	t.ExecuteTemplate(w, "templateactionseleksi.gohtml", map[string]interface{}{
// // 		"Title": "Belajar Seleksi",
// // 		// "Name":  "Rayandra Wandi Marselana",
// // 	})
// // }

// // func TestTemplateSeleksi(t *testing.T) {
// // 	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// // 	rec := httptest.NewRecorder()

// // 	TemplateActionSeleksi(rec, req)

// // 	res := rec.Result()
// // 	body, _ := io.ReadAll(res.Body)
// // 	fmt.Println(string(body))
// // }

// // perbandingan
// // func TemplateActionPerbandingan(w http.ResponseWriter, r *http.Request) {
// // 	t := template.Must(template.ParseFiles("./template/perbandingan.gohtml"))

// // 	t.ExecuteTemplate(w, "perbandingan.gohtml", map[string]interface{}{
// // 		"Title": "Belajar Perbandingan",
// // 		"Value": 45,
// // 	})
// // }

// // func TestTemplatePerbandingan(t *testing.T) {
// // 	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// // 	rec := httptest.NewRecorder()

// // 	TemplateActionPerbandingan(rec, req)

// // 	res := rec.Result()
// // 	body, _ := io.ReadAll(res.Body)
// // 	fmt.Println(string(body))
// // }

// // Range
// // biasanya untuk menampilkan tabel
// // func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
// // 	t := template.Must(template.ParseFiles("./template/range.gohtml"))

// // 	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
// // 		"Hobbi": []string{ //membuat map
// // 			"Gaming",
// // 			"Coding",
// // 			"Sleeping",
// // 		},
// // 	})
// // }

// // func TestTemplatePerbandingan(t *testing.T) {
// // 	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// // 	rec := httptest.NewRecorder()

// // 	TemplateActionRange(rec, req)

// // 	res := rec.Result()
// // 	body, _ := io.ReadAll(res.Body)
// // 	fmt.Println(string(body))
// // }

// // With
// func TemplateActionwith(w http.ResponseWriter, r *http.Request) {
// 	t := template.Must(template.ParseFiles("./template/with.gohtml"))

// 	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
// 		"Title": "Belajar With",
// 		"Name":  "Rayandra Wandi",
// 		"Addres": map[string]interface{}{
// 			"Addres": "Cilengkrang",
// 			"City":   "Bandung",
// 		},
// 	})
// }

// func TestTemplatePerbandingan(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// 	rec := httptest.NewRecorder()

// 	TemplateActionwith(rec, req)

// 	res := rec.Result()
// 	body, _ := io.ReadAll(res.Body)
// 	fmt.Println(string(body))
// }

// ================================================== Template Layout ==================================================
// mengekport header dan footer
// func TemplateLayout(w http.ResponseWriter, r *http.Request) {
// 	t := template.Must(template.ParseFiles( //jarang menggunakan parsefile lebih sering menggunakan embed dan parseglob
// 		"./template/header.gohtml", "./template/body.gohtml", "./template/footer.gohtml", //wajib menyebutkan semua yang akan dipanggil
// 	))

// 	t.ExecuteTemplate(w, "body.gohtml", map[string]interface{}{
// 		"Title": "Belajar Export HEader Footer",
// 		"Name":  "Rayandra Wandi",
// 	})
// }

// func TestTemplateLayout(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// 	rec := httptest.NewRecorder()

// 	TemplateLayout(rec, req)

// 	res := rec.Result()
// 	body, _ := io.ReadAll(res.Body)
// 	fmt.Println(string(body))
// }

// ================================================== Template FUnction ==================================================
// type MyPage struct {
// 	Name string
// }

// func (mypage MyPage) SayHello(name string) string {
// 	return "Hello " + name + ", My Name " + mypage.Name
// }

// // func TemplateFunction(w http.ResponseWriter, r *http.Request) {
// // 	t := template.Must(template.New("Function").Parse(`{{ .SayHello "Wandi"}}`))

// // 	t.ExecuteTemplate(w, "Function", MyPage{
// // 		Name: "Rayandra Wandi Marselana",
// // 	})
// // }

// // func TestTemplateFUnction(t *testing.T) {
// // 	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// // 	rec := httptest.NewRecorder()

// // 	TemplateFunction(rec, req)

// // 	res := rec.Result()
// // 	body, _ := io.ReadAll(res.Body)
// // 	fmt.Println(string(body))
// // }

// // Global Function
// func TemplateGlobalFunction(w http.ResponseWriter, r *http.Request) {
// 	// t := template.Must(template.New("Function").Parse(`{{ len "Wandi"}}`)) //len = menghitung panjang //global function
// 	//membuat function sendiri
// 	t := template.New("Function") //harus di registrasikan
// 	t.Funcs(map[string]interface{}{
// 		"Import": func(value string) string { //Import = key, membuat function sekaligus
// 			return strings.ToUpper(value) //merubah semua menjadi huruf besar
// 		},
// 	})

// 	t = template.Must(t.Parse("{{ Import .Name}}"))

// 	t.ExecuteTemplate(w, "Function", MyPage{
// 		Name: "Rayandra Wandi Marselana",
// 	})
// }

// func TestTemplateGlobalFunction(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// 	rec := httptest.NewRecorder()

// 	TemplateGlobalFunction(rec, req)

// 	res := rec.Result()
// 	body, _ := io.ReadAll(res.Body)
// 	fmt.Println(string(body))
// }

// ================================================== Template Caching ==================================================
// // agar tidak memparse file selalu cukup sekali saja
// //
// //go:embed template/*.gohtml
// var templates embed.FS

// var alltemplate = template.Must(template.ParseFS(templates, "template/*.gohtml")) //dibuatkan saja untuk global //dipanggil akan lebih cepat di server

// func TemplateChaching(w http.ResponseWriter, r *http.Request) {
// 	alltemplate.ExecuteTemplate(w, "simple.gohtml", "Hello Html Golang")
// }

// func TestTemplateChaching(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/", nil)
// 	rec := httptest.NewRecorder()

// 	TemplateChaching(rec, req)

// 	res := rec.Result()
// 	body, _ := io.ReadAll(res.Body)
// 	fmt.Println(string(body))
// }

// ================================================== XSS (Cross Site Scripting) ==================================================
//TONTON VIDIONYA DI UDEMY

// ================================================== Redirect ==================================================
// func RedirectTo(w http.ResponseWriter, r *http.Request) { //Hasil Akhir
// 	fmt.Fprint(w, "Hello Redirect")
// }

// func RedirectForm(w http.ResponseWriter, r *http.Request) { // melakukan redirect ke redirectto
// 	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect) //"/redirect-to" = path yang dituju
// }

// func RedirectOut(w http.ResponseWriter, r *http.Request) { // melakukan redirect ke redirectto
// 	http.Redirect(w, r, "https://www.youtube.com/", http.StatusTemporaryRedirect) //"/redirect-to" = path yang dituju
// }

// func TestRedirect(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/redirect-form", RedirectForm)
// 	mux.HandleFunc("/redirect-to", RedirectTo)
// 	mux.HandleFunc("/redirect-out", RedirectOut)

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: mux,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ================================================== Upload File ==================================================
// func UploadForm(w http.ResponseWriter, r *http.Request) {
// 	myTemplates := template.Must(template.ParseFiles("./template/upload.gohtml"))
// 	myTemplates.ExecuteTemplate(w, "upload.gohtml", nil)
// }

// func Upload(w http.ResponseWriter, r *http.Request) {
// 	myTemplates := template.Must(template.ParseFiles("./template/uploadsuccess.gohtml"))
// 	// r.ParseMultipartForm(32 << 20)//ukurannya 32 mb kalau di ganti bisa
// 	file, fileHandler, err := r.FormFile("file") //mengambil file
// 	if err != nil {
// 		panic(err)
// 	}

// 	fileDestination, err := os.Create("./resource/" + fileHandler.Filename) //tempat untuk menyimpan file
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = io.Copy(fileDestination, file) //mengcopy dari "file" dan di simpan di "filedestination"
// 	if err != nil {
// 		panic(err)
// 	}

// 	name := r.PostFormValue("name") //mengambil data name pada upload.gohtml
// 	myTemplates.ExecuteTemplate(w, "uploadsuccess.gohtml", map[string]interface{}{
// 		"Name": name,
// 		"File": "/static/" + fileHandler.Filename,
// 	})

// }

// func TestUpload(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/upload-form", UploadForm)
// 	mux.HandleFunc("/upload", Upload)
// 	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resource"))))

// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: mux,
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// //go:embed resource/Harley Quin.jpg
// var uploadFileTest []byte

// func TestUploadFile(t *testing.T) {
// 	body := new(bytes.Buffer) //tempan menyimpan

// 	writer := multipart.NewWriter(body) //mengikuti menuliskan format multipart
// 	writer.WriteField("name", "Rayandra Wandi")
// 	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.png") //ada return error
// 	file.Write(uploadFileTest)
// 	writer.Close()//writer di close

// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
// 	request.Header.Set("Content-Type", writer.FormDataContentType())
// 	recorder := httptest.NewRecorder()

// 	Upload(recorder, request)

// 	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
// 	fmt.Println(string(bodyResponse))
// }

// ================================================== Download File ==================================================
// func DownloadFile(writer http.ResponseWriter, request *http.Request) {
// 	file := request.URL.Query().Get("file") //mengambil url file

// 	if file == "" { //jika tidak ada file
// 		writer.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(writer, "Bad Request")
// 		return
// 	}

// 	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
// 	http.ServeFile(writer, request, "./resource/"+file)
// }

// func TestDownloadFile(t *testing.T) {
// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: http.HandlerFunc(DownloadFile),
// 	}

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ================================================== Middle Ware ==================================================
type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler") //bisa melakukan sesuatu perintah program atau apapun sebelum handler dijalankan
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler") //bisa melakukan sesuatu perintah program atau apapun setelah handler dijalankan
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		panic("Ups")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
