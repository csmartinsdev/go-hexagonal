package database_test

import (
	"database/sql"
	"go-hexagonal/src/adapters/database"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var Database *sql.DB

func setUp() {
	Database, _ = sql.Open("sqlite3", ":memory:")
	createTable(Database)
	createProduct(Database)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (id varchar(255), name varchar(255), price float, status varchar(255));`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES ('abc', 'produto 1', 0.0, 'disabled');`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer Database.Close()
	productDb := database.NewProductDb(Database)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "produto 1", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
