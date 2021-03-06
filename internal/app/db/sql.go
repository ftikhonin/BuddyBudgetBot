package dbmanager

const (
	dbName             = "sqlite-database.db"
	createAccountTable = `CREATE TABLE Account (
		"ChatID" integer NOT NULL PRIMARY KEY,
		"Balance" REAL,
		"InitialDate" TEXT,
		"CurrencyID" integer
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
	getList              = `SELECT ID
								  ,ChatID
								  ,OperationTypeID
								  ,Amount
								  ,CategoryID
								  ,Comment
								  ,substr(Moment,1,19) AS Moment
							  FROM Operation 
							 WHERE ChatID = %d 
							   AND substr(Moment,1,4) || substr(Moment,6,2) || substr(Moment,9,2) >= "%s"`
)
