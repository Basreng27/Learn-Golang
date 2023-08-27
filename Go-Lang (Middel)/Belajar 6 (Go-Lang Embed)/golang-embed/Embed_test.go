package golang_embed

import (
	"embed"
	// _ "embed" //harus di deklarasikan dulu
	"fmt"
	"testing"
)

// ================================================== Embed File String ==================================================
// //go:embed version.txt
// var version string

// func TestString(t *testing.T) {
// 	fmt.Println(version) //menampilkan isi pada file version.txt
// }

// ================================================== Embed File ke []Byte ==================================================
// //go:embed rayandra.jpg
// var logo []byte

// func TestByte(t *testing.T) {
// 	err := ioutil.WriteFile("gambar_next.jpg", logo, fs.ModePerm) //mengcopy gambar //audio vidio dll
// 	if err != nil {
// 		panic(err)
// 	}
// }

// ================================================== Embed Multiple Files ==================================================
// // go :embed file/a.txt
// // go :embed file/b.txt
// // go :embed file/c.txt
// var multiple embed.FS //harus seperti ini

// func TestEmbedMultiple(t *testing.T) {
// 	a, _ := multiple.ReadFile("file/a.txt") //memangggil
// 	fmt.Println(string(a))                  //menampilkan

// 	b, _ := multiple.ReadFile("file/b.txt") //memangggil
// 	fmt.Println(string(b))                  //menampilkan

// 	c, _ := multiple.ReadFile("file/c.txt") //memangggil
// 	fmt.Println(string(c))                  //menampilkan
// }

// ================================================== Path Matcher ==================================================
// go :embed file/*.txt
var path embed.FS //harus seperti ini

func TestMatcher(t *testing.T) {
	dir, _ := path.ReadDir("file")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			contene, _ := path.ReadFile("file/" + entry.Name())
			fmt.Println("Content:", string(contene))
		}
	}
}
