package dbWorker

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/jinzhu/gorm"
)

func TestDbListener(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	result := []string{"id", "name", "counter"}
	mock.ExpectQuery("SELECT \\* FROM `Users`").WillReturnRows(sqlmock.NewRows(result).AddRow(1, "Jack", 2))
	mock.ExpectExec("UPDATE `Users`").WithArgs(3, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	gormDB, _ := gorm.Open("mysql", db)
	DbListener(gormDB.LogMode(true))

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
