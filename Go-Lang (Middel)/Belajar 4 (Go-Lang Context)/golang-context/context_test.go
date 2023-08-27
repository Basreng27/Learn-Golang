package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// ================================================== Context ==================================================
// func TestContext(t *testing.T) {
// 	background := context.Background() //membuat context background
// 	fmt.Println(background)

// 	todo := context.TODO() //membuat context todo
// 	fmt.Println(todo)
// }

// ================================================== Parent and Child Context ==================================================

// ================================================== Context with Value ==================================================
// // menambahkan data dan menampilkan data
// func TestContextWithValue(t *testing.T) {
// 	contextA := context.Background() //membuat context/parent dari semua context

// 	contextB := context.WithValue(contextA, "b", "B") //membuat child context yang parentnya A
// 	contextC := context.WithValue(contextA, "c", "C") //membuat child context yang parentnya A

// 	contextD := context.WithValue(contextB, "d", "D") //membuat child context yang parentnya B
// 	contextE := context.WithValue(contextB, "e", "E") //membuat child context yang parentnya B

// 	contextF := context.WithValue(contextC, "f", "F") //membuat child context yang parentnya F

// 	fmt.Println(contextA)
// 	fmt.Println(contextB)
// 	fmt.Println(contextC)
// 	fmt.Println(contextD)
// 	fmt.Println(contextE)
// 	fmt.Println(contextF)

// 	//untuk mendapatkan isi atau value context
// 	fmt.Println(contextF.Value("f")) //dapat isi
// 	fmt.Println(contextF.Value("c")) //mengambil dari milik parent
// 	fmt.Println(contextF.Value("b")) //tidak bisa karena beda parent
// 	fmt.Println(contextA.Value("b")) //tidak bisa mengambil data child
// }

// ================================================== Context with Cancel ==================================================
// // pengulangan yang tidak pernah berhenti looping leak
// // func CreateCounter() chan int {
// // 	destination := make(chan int)
// // 	go func() {
// // 		defer close(destination)
// // 		counter := 1
// // 		for {
// // 			destination <- counter
// // 			counter++
// // 		}
// // 	}()

// // 	return destination
// // }

// // func TestWithCancel(t *testing.T) {
// // 	fmt.Println("Total Goroutine : ", runtime.NumGoroutine()) //awal 2

// // 	destination := CreateCounter()
// // 	for n := range destination {
// // 		fmt.Println("Counter", n)

// // 		if n == 10 {
// // 			break
// // 		}
// // 	}

// // 	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())//jadi 3
// // }

// // contex with cancel
// func CreateCounter(ctx context.Context) chan int {
// 	destination := make(chan int)

// 	go func() {
// 		defer close(destination)
// 		counter := 1
// 		for {
// 			select { //mengecek apakah sudah dibatalkan atau belum
// 			case <-ctx.Done(): //jika data sudah kosong maka akan retun nill
// 				return
// 			default: //jika msh ada data
// 				destination <- counter
// 				counter++
// 			}
// 		}
// 	}()

// 	return destination
// }

// func TestWithCancel(t *testing.T) {
// 	fmt.Println("Total Goroutine : ", runtime.NumGoroutine()) //awal 2
// 	parent := context.Background()                            //membuat parent context
// 	ctx, cancel := context.WithCancel(parent)                 //untuk cancel

// 	destination := CreateCounter(ctx)
// 	for n := range destination {
// 		fmt.Println("Counter", n)

// 		if n == 10 {
// 			break
// 		}
// 	}

// 	cancel() //mengcancel dan mengirim data untuk menghentikan goroutine ke context
// 	time.Sleep(5 * time.Second)

// 	fmt.Println("Total Goroutine : ", runtime.NumGoroutine()) //jadi 2
// }

// ================================================== Context with Timeout ==================================================
// func CreateCounter(ctx context.Context) chan int {
// 	destination := make(chan int)

// 	go func() {
// 		defer close(destination)
// 		counter := 1
// 		for {
// 			select { //mengecek apakah sudah dibatalkan atau belum
// 			case <-ctx.Done(): //jika data sudah kosong maka akan retun nill
// 				return
// 			default: //jika msh ada data
// 				destination <- counter
// 				counter++
// 				time.Sleep(1 * time.Second) //simulasi slow
// 			}
// 		}
// 	}()

// 	return destination
// }

// func TestWithTimeout(t *testing.T) {
// 	fmt.Println("Total Goroutine : ", runtime.NumGoroutine()) //awal 2
// 	parent := context.Background()                            //membuat parent context
// 	ctx, cancel := context.WithTimeout(parent, 5*time.Second) //untuk cancel with timeout /jika 5 detik proses belum selesai maka akan dihentikan
// 	defer cancel()                                            //mengcancel dan mengirim data untuk menghentikan goroutine ke context

// 	destination := CreateCounter(ctx)
// 	for n := range destination {
// 		fmt.Println("Counter", n)

// 		// if n == 10 {
// 		// 	break
// 		// }
// 	}

// 	// time.Sleep(5 * time.Second)

// 	fmt.Println("Total Goroutine : ", runtime.NumGoroutine()) //jadi 2
// }

// ================================================== Context with Deadline ==================================================
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select { //mengecek apakah sudah dibatalkan atau belum
			case <-ctx.Done(): //jika data sudah kosong maka akan retun nill
				return
			default: //jika msh ada data
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) //simulasi slow
			}
		}
	}()

	return destination
}

func TestWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())                  //awal 2
	parent := context.Background()                                             //membuat parent context
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(7*time.Second)) //untuk cancel with timeout /jika 5 detik proses belum selesai maka akan dihentikan
	defer cancel()                                                             //mengcancel dan mengirim data untuk menghentikan goroutine ke context

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)

		// if n == 10 {
		// 	break
		// }
	}

	// time.Sleep(5 * time.Second)

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine()) //jadi 2
}
