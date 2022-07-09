package dbmanager

const (
	dbName             = "sqlite-database.db"
	createAccountTable = `CREATE TABLE Account (
		"ChatID" integer NOT NULL PRIMARY KEY,
		"Balance" REAL,
		"InitialDate" TEXT,
		"CurrencyID" integer
	);`
	createChatTable = `CREATE TABLE Chat (
		"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"FirstName" TEXT,
		"LastName" TEXT,
		"Username" TEXT,
		"Moment" TEXT
		);`
	createOperationTable = `CREATE TABLE Operation (
			"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"ChatID" integer,
			"OperationTypeID" integer,
			"Amount" REAL,
			"CategoryID" integer,
			"Comment" TEXT,
			"Moment" TEXT
		  );`
	createCurrencyTable = `CREATE TABLE Currency (
							"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
							"Name" TEXT,
							"ShortName" TEXT,
							"Label" TEXT
						);`
	createOperationTypeTable = `CREATE TABLE OperationType (
							"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
							"Name" TEXT,
							"ShortName" TEXT
						  );`
	createCategoryTable = `CREATE TABLE Category (
							"ID" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
							"Name" TEXT,
							"ShortName" TEXT,
							"Label" TEXT
							);`

	insertCurrency             = `INSERT INTO Currency(Name, ShortName, Label) VALUES ("United States dollar", "USD", "$");`
	insertOperationTypeExpense = `INSERT INTO OperationType(Name, ShortName) VALUES ("Expense", "-");`
	insertOperationTypeIncome  = `INSERT INTO OperationType(Name, ShortName) VALUES ("Income", "+");`
	insertOperation            = `INSERT INTO Operation(ChatID, OperationTypeID, Amount, CategoryID, Comment, Moment) 
	VALUES ("%d", "%d", "%g","%d","%s","%s");`
	insertAccount = `INSERT INTO Account(ChatID, Balance, InitialDate, CurrencyID ) 
	VALUES ("%d", "%g", "%s","%d");`
	updateAccountBalance = `UPDATE Account SET Balance = Balance + "%g" WHERE ChatID = %d`
	getBalance           = `SELECT Balance as result FROM Account WHERE ChatID =%d LIMIT 1;`
	getCurrency          = `SELECT * FROM Currency`
)
