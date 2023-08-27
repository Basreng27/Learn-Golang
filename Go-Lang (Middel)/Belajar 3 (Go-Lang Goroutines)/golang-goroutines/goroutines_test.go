//================================================== Goroutines ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func HelloWold() { //membuat function
// 	fmt.Println("Hello World")
// }

// func TestCreateGoroutines(t *testing.T) {
// 	go HelloWold()     //memanggil goroutins //tidak cocok digunakan oleh function yang mengembalikan value atau return
// 	fmt.Println("Ups") //melihat functionmana

// 	time.Sleep(1 * time.Second) //agar tidak mati menunggu code program diatas //harus ada untuk menunggu semua program dijalankan
// }

// ================================================== Goroutines Sangat Ringan ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func DisplayNumber(number int) { //membuat function
// 	fmt.Println("Display", number)
// }

// func TestNumberGoroutines(t *testing.T) {
// 	for i := 0; i < 1000; i++ { //pengulangan 1 miliar kali
// 		go DisplayNumber(i)
// 	}

// 	time.Sleep(10 * time.Second) //agar tidak mati menunggu code program diatas //harus ada untuk menunggu semua program dijalankan
// }

// ================================================== Chanel ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestCreateChanel(t *testing.T) {
// 	chanel := make(chan string) //membuat chanel dengan tipedata string

// 	go func() { //membuat annonimus function goroutines
// 		time.Sleep(2 * time.Second)
// 		chanel <- "Rayandra Wandi Marselana" //mengirim data nama kedalam chanel
// 	}()

// 	data := <-chanel
// 	fmt.Println(data)
// 	// chanel <- "Rayandra Wandi Marselana" //mengirim data ke chanel
// 	// data := <-chanel                     //menerima data dichanel dan dimasukan ke variable baru "data"
// 	// fmt.Println(<-chanel)                //menerima / mengirim data chanel kedalam function

// 	// defer close(chanel)//mau eror atau tidak dipaksa di close
// 	close(chanel) //untuk menutup data chanel jika sudah tidak di pakai
// }

// ================================================== Chanel Sebagai Parameter ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func GiveMeResponse(channel chan string) { //membuat parameter yang parameternya channel
// 	time.Sleep(2 * time.Second)           //sleep
// 	channel <- "Rayandra Wandi Marselana" //isi dari channel
// }

// func TestChanelAsParameter(t *testing.T) { //tetsting
// 	channel := make(chan string) //membuat variable channel

// 	go GiveMeResponse(channel) //goroutines

// 	data := <-channel //mengisi variable data dengan channel
// 	fmt.Println(data)
// 	close(channel)

// 	time.Sleep(5 * time.Second)
// }

// ================================================== Chanel In dan Out ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// // perbedaannya di panahnya. jika panahnya di hilangkan bisa mengirim dan menerima
// func OnlyIn(channel chan<- string) { //mengirim data saja
// 	time.Sleep(2 * time.Second)
// 	channel <- "Rayandra Wandi Marselana"
// }

// func OnlyOut(channel <-chan string) { //menerima data
// 	data := <-channel
// 	fmt.Println(data)
// }

// func TestInOutChannel(t *testing.T) {
// 	channel := make(chan string)

// 	go OnlyIn(channel)
// 	go OnlyOut(channel)

// 	time.Sleep(2 * time.Second)
// 	close(channel)
// }

// ================================================== Buffered Channel ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestBufferd(t *testing.T) {
// 	channel := make(chan string, 3) //membuat channel yang bisa mengirim lebih dari 1 data
// 	defer close(channel)

// 	go func() {
// 		channel <- "Rayandra"  //jika lebih dari 3 isinya maka eror "dead lock"
// 		channel <- "Wandi"     //jika lebih dari 3 isinya maka eror "dead lock"
// 		channel <- "Marselana" //jika lebih dari 3 isinya maka eror "dead lock"
// 	}()

// 	go func() {
// 		fmt.Println(<-channel) //melihat data channel
// 		fmt.Println(<-channel) //melihat data channel
// 		fmt.Println(<-channel) //melihat data channel
// 	}()

// 	time.Sleep(2 * time.Second)
// 	fmt.Println(cap(channel)) //melihat panjang buffer
// 	fmt.Println(len(channel)) //melihat jumlah data di dalam buffer
// }

// ================================================== Range Channel ==================================================
// // menerima data lebih dari 1 channel yang tidak diketahui
// package golang_goroutines

// import (
// 	"fmt"
// 	"strconv"
// 	"testing"
// )

// func TestRangeChannel(t *testing.T) {
// 	channel := make(chan string) //membuat channel

// 	go func() { //function anonimus
// 		for i := 0; i < 10; i++ {
// 			channel <- "Perulangan ke-" + strconv.Itoa(i) //convert sting ke int atau sebaliknya
// 		}
// 		close(channel)
// 	}()

// 	for data := range channel { //mengulangi sesuai panjang channel
// 		fmt.Println("Menerima data ", data) //menampilkan isi dari variable data
// 	}

// 	fmt.Println("Selesai")
// }

// ================================================== Select Channel ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func GiveMeResponse(channel chan string) { //membuat parameter yang parameternya channel
// 	time.Sleep(2 * time.Second)           //sleep
// 	channel <- "Rayandra Wandi Marselana" //isi dari channel
// }

// func TestSelectChannel(t *testing.T) {
// 	channel1 := make(chan string) //membuat channel
// 	channel2 := make(chan string)
// 	defer close(channel1) //megclose channel
// 	defer close(channel2)

// 	go GiveMeResponse(channel1) //goroutine channel
// 	go GiveMeResponse(channel2)

// 	counter := 0

// 	for { //perulangan
// 		select { //select
// 		case data := <-channel1: //data channel 1
// 			fmt.Println("Data Dari Channel 1 : ", data) //jika ya maka masuk kesini
// 			counter++
// 		case data := <-channel2:
// 			fmt.Println("Data Dari Channel 2 : ", data)
// 			counter++
// 		}

// 		if counter == 2 { //jika counter isinya 2 maka akan menjalankan perintah break
// 			break
// 		}
// 	}
// }

// ================================================== Default Select ==================================================
// package golanggoroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func GiveMeResponse(channel chan string) { //membuat parameter yang parameternya channel
// 	time.Sleep(2 * time.Second)           //sleep
// 	channel <- "Rayandra Wandi Marselana" //isi dari channel
// }

// func TestSelectChannel(t *testing.T) {
// 	channel1 := make(chan string) //membuat channel
// 	channel2 := make(chan string)
// 	defer close(channel1) //megclose channel
// 	defer close(channel2)

// 	go GiveMeResponse(channel1) //goroutine channel
// 	go GiveMeResponse(channel2)

// 	counter := 0

// 	for { //perulangan
// 		select { //select
// 		case data := <-channel1: //data channel 1
// 			fmt.Println("Data Dari Channel 1 : ", data) //jika ya maka masuk kesini
// 			counter++
// 		case data := <-channel2:
// 			fmt.Println("Data Dari Channel 2 : ", data)
// 			counter++
// 		default ://hanya menambahkan ini
// 		fmt.Println("Menunggu data")//jika data tidak ada maka akan masuk kesini
// 		}

// 		if counter == 2 { //jika counter isinya 2 maka akan menjalankan perintah break
// 			break
// 		}
// 	}
// }

// ================================================== Race Condition ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestRaceCondition(t *testing.T) {
// 	//code untuk cek race condition
// 	var x = 0
// 	for i := 0; i <= 1000; i++ {
// 		go func() {
// 			for j := 1; j <= 100; j++ {
// 				x = x + 1
// 			}
// 		}()
// 	}

// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Counter ", x)
// }

// ================================================== sync.Mutex ==================================================
// //lebih aman menggunakan mutex akan tetapi lebih lambat
// package golang_goroutines

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"
// )

// func TestMutex(t *testing.T) {
// 	var x = 0
// 	var mutex sync.Mutex //membuat variable mutex

// 	for i := 1; i <= 1000; i++ {
// 		go func() {
// 			for j := 1; j <= 100; j++ {
// 				mutex.Lock()   //melock data dulu
// 				x = x + 1      //sebelum merubah x lock dulu kemudain di unlock
// 				mutex.Unlock() //kemudian diunlock
// 			}
// 		}()
// 	}

// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Counter ", x)
// }

// ================================================== sync.RWMutex ==================================================
// // yang akan di akses beberapa goroutines
// package golang_goroutines

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"
// )

// type BankAccount struct { //membuat struct
// 	RWMutex sync.RWMutex //membuat variable RWMutex
// 	Balance int
// }

// func (account *BankAccount) AddBalance(amount int) { //function/method
// 	account.RWMutex.Lock() //write
// 	account.Balance = account.Balance + amount
// 	account.RWMutex.Unlock()
// }

// func (account *BankAccount) GetBalance() int {
// 	account.RWMutex.RLock() //read locking
// 	balance := account.Balance
// 	account.RWMutex.RUnlock() //read unlocking

// 	return balance
// }

// func TestRWMutex(t *testing.T) {
// 	account := BankAccount{}

// 	for i := 0; i <= 100; i++ {
// 		go func() {
// 			for j := 1; j <= 100; j++ {
// 				account.AddBalance(1)
// 				fmt.Println(account.GetBalance())
// 			}
// 		}()
// 	}

// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Total Balance ", account.GetBalance())
// }

// ================================================== Dead Lock ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"
// )

// type UserBalance struct { //membuat struct
// 	sync.Mutex
// 	name    string
// 	Balance int
// }

// func (user *UserBalance) Lock() { //membuat lock
// 	user.Mutex.Lock()
// }

// func (user *UserBalance) Unlock() { //membuat unlock
// 	user.Mutex.Unlock()
// }

// func (user *UserBalance) Change(amount int) {
// 	user.Balance = user.Balance + amount
// }

// func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
// 	user1.Lock()
// 	fmt.Println("Lock", user1.name)
// 	user1.Change(-amount)

// 	time.Sleep(1 * time.Second)

// 	user2.Lock()
// 	fmt.Println("Lock", user2.name)
// 	user2.Change(amount)

// 	time.Sleep(1 * time.Second)

// 	user1.Unlock()
// 	user2.Unlock()
// }

// func TestDeadLock(t *testing.T) {
// 	user1 := UserBalance{
// 		name: "Rayandra",
// 	}

// 	user2 := UserBalance{
// 		name: "wandi",
// 	}

// 	go Transfer(&user1, &user2, 27) //mengirimkan isinya kedalam method Transfer
// 	go Transfer(&user1, &user2, 22)

// 	time.Sleep(5 * time.Second)

// 	fmt.Println("Nama : ", user1.name, " Balance : ", user1.Balance)
// 	fmt.Println("Nama : ", user2.name, " Balance : ", user2.Balance)
// }

// ================================================== sync.WaitGroup ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"
// )

// // wait code
// func RunAsynchonus(group *sync.WaitGroup) { //membuat parameter waitgroup dengan pointer
// 	defer group.Done()

// 	group.Add(1) //menjalankan go routine

// 	fmt.Println("Hello")
// 	time.Sleep(1 * time.Second)
// }

// func TestWaitGroup(t *testing.T) {
// 	group := &sync.WaitGroup{}

// 	for i := 0; i < 100; i++ {
// 		go RunAsynchonus(group)
// 	}

// 	group.Wait()
// 	fmt.Println("Selesai")
// }

// ================================================== sync.Once ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// )

// var counter = 0

// func OnlyOnce() {
// 	counter++
// }

// func TestOnce(t *testing.T) {
// 	var once sync.Once //membuat variable once
// 	var group sync.WaitGroup

// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			group.Add(1)
// 			once.Do(OnlyOnce) //memanggil function yang bisa di akses 1x
// 			group.Done()
// 		}()
// 	}

// 	group.Wait()
// 	fmt.Println(counter)
// }

// ================================================== sync.Pool ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"
// )

// func TestPool(t *testing.T) {
// 	// var pool sync.Pool//manual

// 	//kode membuat data pool otomatis
// 	var pool = sync.Pool{//kembaliannya interface kosong
// 		New: funct() interface{}{
// 			return "New"
// 		}
// 	}

// 	pool.Put("Rayandra")  //memasukan data menggunakan put
// 	pool.Put("Wandi")     //memasukan data menggunakan put
// 	pool.Put("Marselana") //memasukan data menggunakan put

// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			data := pool.Get() //memanggil data menggunakan det
// 			fmt.Println(data)  //mengambil dan menampilkan
// 			time.Sleep(1 * time.Second)
// 			pool.Put(data) //jika selesai mengambil data maka kembalikan lagi datanya
// 		}()

// 		time.Sleep(3 * time.Second)
// 	}
// }

// ================================================== sync.Map ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"
// )

// func TestSyncMao(t *testing.T) {
// 	var data = sync.Map{}
// 	var addMap = func(value int) {
// 		data.Store(value, value)
// 	}

// 	for i := 0; i < 100; i++ {
// 		go addMap(i)
// 	}

// 	time.Sleep(3 * time.Second)
// 	data.Range(func(key, value interface{}) bool { //lanjut ke selanjutnya jika true false berhenti
// 		fmt.Println(key, " : ", value)
// 		return true
// 	})
// }

// ================================================== sync.Cond ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"sync"
// 	"testing"
// 	"time"
// )

// var cond = sync.NewCond(&sync.Mutex{}) //membuat variable cond diisi cond baru mutex tapi pointernya
// var group = &sync.WaitGroup{}

// func WaitCondition(value int) {
// 	cond.L.Lock() //lock
// 	cond.Wait()
// 	fmt.Println("Done", value)
// 	cond.L.Unlock()
// 	group.Done()
// }

// func TestCond(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		group.Add(1)
// 		go WaitCondition(i)
// 	}

// 	//signal
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(1 * time.Second)
// 			cond.Signal() //memberi sinyal bahwa melanjutkan goroutine satu satu
// 		}
// 	}()

// 	//broadcast
// 	// go func() {
// 	// 	time.Sleep(1 * time.Second)
// 	// 	cond.Broadcast() //memberi sinyal bahwa melanjutkan goroutine semua

// 	// }()

// 	group.Wait()
// }

// ================================================== Atomic ==================================================
// package golang_goroutines

// ================================================== time.Timer ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// // time.Newtimer
// func TestTimer(t *testing.T) {
// 	timer := time.NewTimer(5 * time.Second) //menunggu selama 5 detik

// 	fmt.Println(time.Now())

// 	time := <-timer.C //c adalah chanel
// 	fmt.Println(time)
// }

// // time.After
// func TestAfter(t *testing.T) {
// 	timer := time.After(5 * time.Second)

// 	time := <-timer //c adalah chanel
// 	fmt.Println(time)
// }

// //time.AfterFunction

// ================================================== time.Timer ==================================================
// package golang_goroutines

// import (
// 	"fmt"
// 	"testing"
// 	"time"
// )

// func TestTicker(t *testing.T) {
// 	ticker := time.NewTicker(1 * time.Second)

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		ticker.Stop()
// 	}()

// 	for tick := range ticker.C { //ulangi sesuai panjang tricker.c atau chanel
// 		fmt.Println(tick)
// 	}
// }

// ================================================== GOMAXPROCS ==================================================
package golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomax(t *testing.T) {
	totalcpu := runtime.NumCPU()
	fmt.Println("CPU", totalcpu)

	totalthread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalthread)

	totalgorountine := runtime.NumGoroutine()
	fmt.Println("Gorountine", totalgorountine)
}
