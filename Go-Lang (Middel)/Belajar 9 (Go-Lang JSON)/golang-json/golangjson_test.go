package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// ================================================== Encode Json ==================================================
func LogJson(data interface{}) { //tipe data adalah interface kosong
	bytes, err := json.Marshal(data) //untuk encode dari tipedata golang ke json
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonMarshalEncode(t *testing.T) {
	LogJson("Wandi")
	LogJson(1)
	LogJson(true)
	LogJson([]string{ //slice
		"Rayandra",
		"Wandi",
		"Marselana",
	})
}

// ================================================== Json Object ==================================================
type Costumer struct { //struct
	Firstname, Middlename, Lastname string
}

func TestJsonObject(t *testing.T) {
	costumer := Costumer{
		Firstname:  "Rayandra",
		Middlename: "Wandi",
		Lastname:   "Marselana",
	}

	bytes, _ := json.Marshal(costumer) //mengkonversikan
	fmt.Println(string(bytes))
}

// ================================================== Json Decode ==================================================
func TestDecodeJson(t *testing.T) {
	jsonReq := `{"Firstname":"Rayandra","Middlename":"Wandi","Lastname":"Marselana"}` //data json yang ingin diubah ke golang kebalikan encode
	jsonByte := []byte(jsonReq)                                                       //di masukan kedalam byte

	costumer := &Costumer{} //atribut atas akan di masukan kedalam struct Costumer
	json.Unmarshal(jsonByte, costumer)

	fmt.Println(costumer)
}

// ================================================== Json Array ==================================================
type Addres struct {
	Street, City, Country string
}

type Pelanggan struct {
	First, Mid, Last string
	Hobbies          []string //tipe data string slice
	Address          []Addres //slice dari addres
}

// encode
func TestJsonArrayEncode(t *testing.T) {
	pelanggan := Pelanggan{
		First: "Rayandra",
		Mid:   "Wandi",
		Last:  "Marselana",
		Hobbies: []string{ //slice
			"Coding",
			"Game",
			"Sleep",
		},
	}

	byte, _ := json.Marshal(pelanggan)
	fmt.Println(string(byte))
}

// decode
func TestJsonArrayDecode(t *testing.T) {
	jsonReq := `{"First":"Rayandra", "Mid":"Wandi", "Last":"Marselana","Hobbies":["Coding","Game","Sleep"]}`

	jsonbyte := []byte(jsonReq)
	stri := &Pelanggan{}

	err := json.Unmarshal(jsonbyte, stri)
	if err != nil {
		panic(err)
	}

	fmt.Println(stri)
}

func TestJsonArrayComplex(t *testing.T) {
	pelanggan := Pelanggan{
		First: "Rayandra",
		Mid:   "Wandi",
		Last:  "Marselana",
		Hobbies: []string{ //slice
			"Coding",
			"Game",
			"Sleep",
		},
		Address: []Addres{{
			Street:  "Cilengkrang",
			City:    "Bandung",
			Country: "Indonesia",
		},
		},
	}

	byte, _ := json.Marshal(pelanggan)
	fmt.Println(string(byte))
}

// ================================================== Json Tag ==================================================
type Product struct {
	Id    string `json:"id"` //merubah yang asalnya Id menjadi id Di json
	Name  string `json:"name"`
	Price int64  `json:price`
}

func TestJsonTag(t *testing.T) {
	//panggil mengunakan json.Marshal
}

// ================================================== Map ==================================================
// lebih baik menggunakan map karena datanya bisa berubah ubah sedangkan struct harus tetap
func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"p0001", "name":"Apple Mac Book Pro", "price": 20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Apple Mac Book Pro",
		"price": 20000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

// ================================================== Decoder ==================================================
type Orang struct {
	Firstname, Middlename, Lastname string
}

func TestStreamDecoder(t *testing.T) {

	reader, _ := os.Open("customer.json") //membaca file .json
	decoder := json.NewDecoder(reader)    //membaca ini atau menggunakan ini

	customer := &Orang{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

// ================================================== Encoder ==================================================
func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json") //menulis ke file langsung
	encoder := json.NewEncoder(writer)         //menggunakan ini

	customer := Orang{
		Firstname:  "Eko",
		Middlename: "Kurniawan",
		Lastname:   "Khannedy",
	}

	encoder.Encode(customer) //maka akan langsung ada file bernama costumerout.json
}
