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
	// 创建模拟数据库连接
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer db.Close()

	// 设置模拟数据库的预期行为
	// mock.ExpectExec("INSERT INTO Ad").WillReturnResult(sqlmock.NewResult(1, 1))

	// 假設的輸入數據
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

	// 检查模拟数据库的期望行为是否满足
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}