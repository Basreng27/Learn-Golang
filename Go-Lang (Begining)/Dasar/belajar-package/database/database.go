package database

var connection string //membuat variable

func init() { //untuk function initialization / function awal di eksekusi
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
