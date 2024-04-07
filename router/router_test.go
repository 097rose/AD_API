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
			req, err := http.NewRequest("GET", "/api/v1/ad", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			w := httptest.NewRecorder()
			got.ServeHTTP(w, req)

			//check ststus 200
			if w.Code != http.StatusOK {
				t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
			}

			
		})
	}
}


func TestSetupRouter_PostRequest(t *testing.T) {
	// test request
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

	req, err := http.NewRequest("POST", "/api/v1/ad", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	w := httptest.NewRecorder()
	router := SetupRouter()
	router.ServeHTTP(w, req)

	// check status 200 
	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", w.Code, http.StatusOK)
	}

	
}
