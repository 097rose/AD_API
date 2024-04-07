package router

import (
	
	"net/http"
	"net/http/httptest"
	// "reflect"
	"testing"
	"bytes"

	"github.com/gin-gonic/gin"
)

func TestSetupRouter_GET(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		{
			name: "Test SetupRouter",
			want: gin.Default(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetupRouter()

			// 创建一个 HTTP 请求
			req, err := http.NewRequest("GET", "/api/v1/ad", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			// 创建一个响应记录器
			w := httptest.NewRecorder()

			// 使用创建的路由处理请求
			got.ServeHTTP(w, req)

			// 检查状态码是否为 200
			if w.Code != http.StatusOK {
				t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
			}

			
		})
	}
}


func TestSetupRouter_PostRequest(t *testing.T) {
	// 准备请求体
	requestBody := []byte(`{
		"title": "AD 405",
		"startAt": "2023-12-10T03:00:00.000Z",
		"endAt": "2024-06-11T16:00:00.000Z",
		"conditions": [
			{
				"ageStart": 10,
				"ageEnd": 39,
				"country": ["TW", "JP"],
				"platform": ["android", "ios"]
			}
		]
	}`)

	// 创建一个新的 HTTP 请求
	req, err := http.NewRequest("POST", "/api/v1/ad", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// 创建一个响应记录器
	w := httptest.NewRecorder()

	// 调用 SetupRouter 处理请求
	router := SetupRouter()
	router.ServeHTTP(w, req)

	// 检查返回的状态码是否为 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", w.Code, http.StatusOK)
	}

	
}
