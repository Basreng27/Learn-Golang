package helper

import "database/sql"

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollbeck := tx.Rollback() //mengembalikan data aatu tidak jadi di input
		PanicIfError(errorRollbeck)
		panic(err)
	} else {
		errorCommit := tx.Commit() //menyetujui di input
		PanicIfError(errorCommit)
	}
}
