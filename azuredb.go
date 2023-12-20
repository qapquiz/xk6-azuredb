package azuredb 

import (
	"database/sql"
	"log"

	"github.com/microsoft/go-mssqldb/azuread"
	// "go.k6.io/k6/js/modules"
)

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/azuredb", new(AzureDB))
}

type AzureDB struct {}
type InsertResult struct {
	ID int64
}

func connectDatabase(connectionString string) (*sql.DB, error) {
	db, err := sql.Open(
	    azuread.DriverName,
	    connectionString,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func (az *AzureDB) Insert(connectionString string, rawQueries []string) ([]int64, error) {
	db, err := connectDatabase(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	insertedIds := []int64 {}

	for _, rawQuery := range rawQueries {
		var insertResult InsertResult
		row := db.QueryRow(rawQuery)
		
		if err := row.Scan(&insertResult.ID); err != nil {
			log.Fatal(err)
		}

		insertedIds = append(insertedIds, insertResult.ID)
	}

	return insertedIds, err
}
