package api

import (
	"Api/model"
	"reflect"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"

	_ "github.com/go-sql-driver/mysql"
)


func TestSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("connect fail：%s", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"title", "startAt", "endAt"}).AddRow("Test AD", "2023-12-10T03:00:00.000Z", "2024-06-11T16:00:00.000Z"))
	result := Select(0, 10, "20", 1, "US", "android")

	
	if len(result.Items) == 1 {
		t.Errorf("error")
	}
	
}
