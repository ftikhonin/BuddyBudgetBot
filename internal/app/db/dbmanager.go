package dbmanager

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	if _, err := os.Stat(dbName); err != nil {
		log.Println("Creating sqlite-database.db...")
		file, err := os.Create(dbName)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		Execute(createAccountTable)
		Execute(createOperationTable)
		Execute(createCurrencyTable)
		Execute(createOperationTypeTable)
		Execute(createCategoryTable)
		Execute(insertCurrency)
		Execute(insertOperationTypeExpense)
		Execute(insertOperationTypeIncome)

		log.Println("Database created")
	}

}

func Execute(statement string) error {

	var db *sql.DB
	var err error
	if _, err := os.Stat(dbName); err == nil {
		db, _ = sql.Open("sqlite3", "./sqlite-database.db")

		log.Println("Executing statement: " + statement)
		preparedStatement, err := db.Prepare(statement)

		if err != nil {
			log.Println(err.Error())
			return err
		}
		_, err = preparedStatement.Exec()

		if err != nil {
			log.Println(err.Error())
			return err
		}

		log.Println("Success")

		defer db.Close()
	}
	return err
}

func ExecuteScalar(statement string) (string, error) {

	var db *sql.DB
	var _ string
	var result string
	var err error
	if _, err = os.Stat(dbName); err == nil {
		db, _ = sql.Open("sqlite3", "./sqlite-database.db")
		defer db.Close()

		log.Println("Executing statement: " + statement)
		err = db.QueryRow(statement).Scan(&result)
		if err != nil {
			return result, err
		}
		log.Println("Execute success: " + statement)

	}
	return result, nil
}

func displayCurrency(db *sql.DB) string {
	row, err := db.Query(getCurrency)
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
		return "Currency: " + Name + " " + ShortName + " " + Label
	}
	return ""
}

func Register(chatID int64, balance float32) {
	Execute(
		fmt.Sprintf(insertAccount, chatID, balance, time.Now(), 1))
}

func GetBalance(chatID int64) (string, error) {

	result, err := ExecuteScalar(
		fmt.Sprintf(getBalance, chatID))
	if err != nil {
		err = errors.New("First, register using the command /req")
		return "", err
	}
	return result, nil

}

func GetList(chatID int64, forDate string) (string, error) {
	var db *sql.DB
	var _ string
	var err error
	if _, err = os.Stat(dbName); err == nil {
		db, _ = sql.Open("sqlite3", "./sqlite-database.db")
		defer db.Close()
		statement := fmt.Sprintf(getList, chatID, forDate)
		log.Println("Executing statement: " + statement)
		row, err := db.Query(statement)
		if err != nil {
			log.Fatal(err)
		}
		defer row.Close()
		var result string = ""
		for row.Next() {
			var ID int
			var ChatID string
			var OperationTypeID int
			var Amount string
			var CategoryID int
			var Comment string
			var Moment string
			row.Scan(&ID, &ChatID, &OperationTypeID, &Amount, &CategoryID, &Comment, &Moment)

			if OperationTypeID == 1 {
				result += "\U00002796   " //:heavy_minus_sign:
			} else {
				result += "\U00002795   " //:heavy_plus_sign:
			}

			layout := "2006-01-02 15:04:05" //magical reference date
			t, _ := time.Parse(layout, Moment)
			result += Amount + " " + "   " + t.Format("02-Jan-2006") + "\n"
		}
		balance, _ := GetBalance(chatID)
		result += "________________________\n" + balance
		return result, nil
	}
	return "", nil

}

func SetIncome(chatID int64, amount float64, comment string) (string, error) {

	err := Execute(
		fmt.Sprintf(insertOperation, chatID, 2, amount, 1, comment, time.Now()))
	if err != nil {
		return "", err
	}

	Execute(
		fmt.Sprintf(updateAccountBalance, amount, chatID))
	if err != nil {
		return "", err
	}

	return GetBalance(chatID)

}

func SetExpense(chatID int64, amount float64, comment string) (string, error) {
	result, _ := ExecuteScalar(fmt.Sprintf(getBalance, chatID))
	currentBalance, _ := strconv.ParseFloat(result, 64)
	if currentBalance < amount {
		return "", errors.New("The amount cannot be more than the current balance")
	}

	err := Execute(
		fmt.Sprintf(insertOperation, chatID, 1, amount, 1, comment, time.Now()))
	if err != nil {
		return "", err
	}

	Execute(
		fmt.Sprintf(updateAccountBalance, -amount, chatID))
	if err != nil {
		return "", err
	}

	return GetBalance(chatID)

}
