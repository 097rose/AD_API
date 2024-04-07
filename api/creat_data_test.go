package api

import (

	"testing"
	"Api/model"
	// "database/sql"
	// "fmt"
	// "log"
	// "reflect"
	// "time"
	"github.com/DATA-DOG/go-sqlmock"

	_ "github.com/go-sql-driver/mysql"
)


func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	//using test data
	testData := model.PostBody{
		Title: "AD 405",
		StartAt: "2023-12-10T03:00:00.000Z",
		EndAt: "2024-06-11T16:00:00.000Z",
		Conditions: []struct{
			AgeStart int `json:"ageStart"`
			AgeEnd int `json:"ageEnd"`
			Country []string `json:"country"`
			Gender string `json:"gender"`
			Platform []string `json:"platform"`
		}{
			{
				AgeStart: 10,
				AgeEnd: 39,
				Country: []string{"TW", "JP"},
				Gender: "M",
				Platform: []string{"android", "ios"},
			},
		},
	}

	Create(testData)
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}