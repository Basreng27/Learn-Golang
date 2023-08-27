// ================================================== Exekusi Perintah Sql ==================================================
package golang_database

// insert table ke database
// func TestInsertDatabase(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	script := "INSERT INTO costumer(id,name) VALUE ('1218015','Adrian Nugraha');" //bisa dicoba insert update, delete
// 	_, err := db.ExecContext(ctx, script)                                         //perintah insert data //jika untuk mengambil data tidak bisa
// 	// _, err := db.ExecContext(ctx, "INSERT INTO costumer(id,name) VALUE ('1218022','Rayandra Wandi Marselana');") //perintah insert data //jika untuk mengambil data tidak bisa
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Insert data berhasil")
// }

// ================================================== Query SQL ==================================================
// // untuk select database
// func TestSelectDatabase(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()
// 	// script := "SELECT * FROM costumer"//Semua data
// 	script := "SELECT * FROM costumer WHERE id='1218022'" //mengambil data tertentu

// 	rows, err := db.QueryContext(ctx, script)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	fmt.Println("Data Berhasil Diambil")

// 	//kode rows untuk menampilkan
// 	for rows.Next() { //mengulangi sebanyak data rows
// 		var id, name string
// 		err := rows.Scan(&id, &name) //melakukan pemanggilan dari field //harus pointer asli atau data asli

// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("")
// 		fmt.Println("Data Diri : ")
// 		fmt.Println("Id	: ", id)      //dipanggil
// 		fmt.Println("Nama 	: ", name) //dipanggil
// 	}
// 	defer db.Close() //jangan lupa di tutup rowsnya
// }

// ================================================== Tipe Data Column ==================================================
// func TestDataColumn(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	// // syntax1 := "INSERT INTO costummer (id,name,email,balance,rating,birth_date,maried) VALUES (1218022,'Rayandra Wandi Marselana','wandirayandra@gmail.com',4000000,5.0,'1999-01-01',true)"
// 	// syntax2 := "INSERT INTO costummer (id,name,email,balance,rating,birth_date,maried) VALUES (1218023,'Raya Lil Ayunda','lilraya@gmail.com',2000000,3.0,'2010-10-10',true)"
// 	// _, err := db.ExecContext(ctx, syntax2)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	var selet string = "SELECT * FROM costummer"
// 	rows, err := db.QueryContext(ctx, selet)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id, name string
// 		var email sql.NullString //untuk bisa data null atau kosong
// 		var balance int32
// 		var rating int64
// 		var birth_date time.Time
// 		var maried bool

// 		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &maried)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("========================================")
// 		fmt.Println("Id	: ", id)
// 		fmt.Println("Name	: ", name)
// 		if email.Valid { //pengecekan ada atau tidak nul
// 			fmt.Println("Email	: ", email.String)
// 		}
// 		fmt.Println("Balance	: ", balance)
// 		fmt.Println("Rating	: ", rating)
// 		fmt.Println("Birth Date	: ", birth_date)
// 		fmt.Println("Married	: ", maried)
// 	}
// 	defer db.Close()
// }

// ================================================== Sql Injection ==================================================
// //Query dengan parameter /biasa di gunakan login

// func TestSqlInjection(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	// username := "admin"
// 	username := "admin `; #" //injection //user merubah sql
// 	password := "admin"

// 	ctx := context.Background()
// 	sqlQuery := "select username from user where username ='" + username + "' and password ='" + password + "' limit 1" //+ = menggabungkan string dengan variable
// 	rows, err := db.QueryContext(ctx, sqlQuery)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		var username string

// 		err := rows.Scan(&username)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("Sukses Login", username)
// 	} else {
// 		fmt.Println("Gagal Login", username)
// 	}
// }

// ================================================== Sql Dengan Parameter ==================================================
// func TestSqlInjection(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	// username := "admin"
// 	username := "admin `; #" //injection //user merubah sql
// 	password := "admin"

// 	ctx := context.Background()
// 	sqlQuery := "select username from user where username = ? and password = ? limit 1" //+ = menggabungkan string dengan variable
// 	rows, err := db.QueryContext(ctx, sqlQuery, username, password)                     //lebih aman menggunakan ini //bisa digunakan di ExeContext
// 	// sqlinput := "insert into user(username, password) values (?,?)"//input
// 	// _, err := db.ExecContext(ctx, sqlinput, username, password)                         //input
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		var username string

// 		err := rows.Scan(&username)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println("Sukses Login", username)
// 	} else {
// 		fmt.Println("Gagal Login", username)
// 	}
// }

// ================================================== Auto Increment ==================================================
// // insert id autoincrement
// func TestAutoIncrement(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	email := "wandirayandra@gmail.com"
// 	comment := "Hallo Gaiz"

// 	ctx := context.Background()
// 	sqlQ := "insert into comment(email,comment) values (?,?)"
// 	result, err := db.ExecContext(ctx, sqlQ, email, comment)
// 	if err != nil {
// 		panic(err)
// 	}

// 	insertId, err := result.LastInsertId() //untuk mengambil id terakhir
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Id Terakhir yang dimasukan : ", insertId)
// }

// ================================================== Prepare Statement ==================================================
// func TestPrepareStatement(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()
// 	stmt, err := db.PrepareContext(ctx, "INSERT INTO comment(email,comment) VALUES (?,?)")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()

// 	//input dengan beberapa data
// 	for i := 0; i < 10; i++ {
// 		email := "wandi" + strconv.Itoa(i) /*Mengganti int menjadi string*/ + "@gmail.com"
// 		comment := "Ini Komen Ke-" + strconv.Itoa(i)

// 		result, err := stmt.ExecContext(ctx, email, comment) //memakai stmt
// 		if err != nil {
// 			panic(err)
// 		}

// 		lastInsert, _ := result.LastInsertId()
// 		fmt.Println("Last Comment Id : ", lastInsert)
// 	}
// }

// ================================================== Database Transaction ==================================================
// func TestXDatabaseTransaction(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	tx, err := db.Begin()
// 	if err != nil {
// 		panic(err)
// 	}

// 	//lakukan transaksi disini
// 	script := "INSERT INTO comment(email,comment) VALUES (?,?)"
// 	for i := 0; i < 10; i++ {
// 		email := "wandi" + strconv.Itoa(i) /*Mengganti int menjadi string*/ + "@gmail.com"
// 		comment := "Ini Komen Ke-" + strconv.Itoa(i)

// 		result, err := tx.ExecContext(ctx, script, email, comment) //tx
// 		if err != nil {
// 			panic(err)
// 		}

// 		lastInsert, _ := result.LastInsertId()
// 		fmt.Println("Last Comment Id : ", lastInsert)
// 	}

// 	er := tx.Rollback() //Membatalkan transaction
// 	// er := tx.Commit() //jika ingin membatalkan bisa gunakan "Rolback()"
// 	if er != nil {
// 		panic(er)
// 	}
// }

// ================================================== Reppository Pattern ==================================================
