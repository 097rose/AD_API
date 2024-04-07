package api

import (
	"Api/model"
	"reflect"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelect(t *testing.T) {
	type args struct {
		offset   int
		limit    int
		age      string
		gender   int
		country  string
		platform string
	}
	tests := []struct {
		name string
		args args
		want model.ResultBody
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(tt.args.offset, tt.args.limit, tt.args.age, tt.args.gender, tt.args.country, tt.args.platform); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelect2(t *testing.T) {
	// 准备测试用例
	tests := []struct {
		offset   int
		limit    int
		age      string
		gender   int
		country  string
		platform string
		want     model.ResultBody
	}{
		{
			offset:   0,
			limit:    10,
			age:      "20",
			gender:   1,
			country:  "TW",
			platform: "android",
			want: model.ResultBody{
				Items: []model.ResultItem{
					{
						Title:  "AD 405",
						StartAt: "2023-12-10T03:00:00.000Z",
						EndAt: "2024-06-11T16:00:00.000Z",
					},
					// 可根据需要添加更多预期结果
				},
			},
		},
		// 可添加更多测试用例
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Select(tt.offset, tt.limit, tt.age, tt.gender, tt.country, tt.platform); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSelect3(t *testing.T) {
	// 创建模拟数据库连接
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("创建模拟数据库连接时发生错误：%s", err)
	}
	defer db.Close()

	// 设置模拟的行为和预期结果
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"title", "startAt", "endAt"}).AddRow("Test AD", "2023-12-10T03:00:00.000Z", "2024-06-11T16:00:00.000Z"))

	// 调用函数
	result := Select(0, 10, "20", 1, "US", "android")

	// 验证结果
	if len(result.Items) == 1 {
		t.Errorf("查询结果数量错误，期望 1 条，实际 %d 条", len(result.Items))
	}
	
}
