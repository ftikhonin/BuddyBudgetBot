package dbmanager

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {

	os.Remove("sqlite-database.db")

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer sqliteDatabase.Close()
	createTable(sqliteDatabase)
	insertDefault(sqliteDatabase)
	displayCurrency(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE Account (
		"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"ChatID" TEXT,
		"Balance" REAL,
		"InitialDate" TEXT,
		"CurrencyID" integer
	  );
	
	  CREATE TABLE Chat (
		"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"FirstName" TEXT,
		"LastName" TEXT,
		"Username" TEXT,
		"Moment" TEXT				
	  );
	
	  CREATE TABLE Operation (
		"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"AccountID" integer,
		"OperationTypeID" integer,
		"Amount" REAL,
		"CategoryID" integer
		"Comment" TEXT
		"Moment" TEXT		
	  );
	
	  CREATE TABLE Currency (
		"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"Name" TEXT,
		"ShortName" TEXT,
		"Label" TEXT		
	  );
	
	  CREATE TABLE OperationType (
		"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"Name" TEXT,
		"ShortName" TEXT	
	  );
	
	  CREATE TABLE Category (
		"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"Name" TEXT,
		"ShortName" TEXT,
		"Label" TEXT		
	  );`

	log.Println("Creating tables...")
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Tables created")
}

func insertDefault(db *sql.DB) {
	log.Println("Inserting default record ...")
	insertStudentSQL := `INSERT INTO Currency(Name, ShortName, Label) VALUES ("United States dollar", "USD", "$");
	INSERT INTO OperationType(Name) VALUES ("Income");
	INSERT INTO OperationType(Name) VALUES ("Expense");
	INSERT INTO Category(Name) VALUES ("Deposits");
	INSERT INTO Category(Name) VALUES ("Salary");
	INSERT INTO Category(Name) VALUES ("Bills");
	INSERT INTO Category(Name) VALUES ("Food");`
	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayCurrency(db *sql.DB) {
	row, err := db.Query("SELECT * FROM Currency")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var Name string
		var ShortName string
		var Label string
		row.Scan(&id, &Name, &ShortName, &Label)
		log.Println("Currency: ", Name, " ", ShortName, " ", Label)
	}
}
