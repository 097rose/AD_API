package router

import (
	"Api/api"
	"Api/model"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// "fmt"
	"strconv"
	"time"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/api/v1/ad", CreateAD)
    r.GET("/api/v1/ad", SelectAD)
    return r
}

func CreateAD(c *gin.Context) {
    var request model.PostBody
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }else{
        startAtTime, err := time.Parse(time.RFC3339, request.StartAt)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "startTime error"})
            return
        }

        endAtTime, err := time.Parse(time.RFC3339, request.EndAt)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "endTime error"})
            return
        }

        if startAtTime.After(endAtTime) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "startTime should before endTime"})
            return
        }

        if request.Title == "" || startAtTime.IsZero() || endAtTime.IsZero() {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Title, StartAt, and EndAt cannot be empty"})
            return
        }
        if request.Conditions[0].AgeEnd == 0{
            request.Conditions[0].AgeEnd = 101
        }
        if request.Conditions[0].AgeEnd < request.Conditions[0].AgeStart{
            c.JSON(http.StatusBadRequest, gin.H{"error": "AgeStart should be smaller than AgeEnd"})
            return
        }
    }
    api.Create(request)
    c.JSON(http.StatusOK, request)
}

func SelectAD(c *gin.Context) {
    
    offsetStr := c.DefaultQuery("offset", "0")
    offset, _ := strconv.Atoi(offsetStr)
    limitStr := c.DefaultQuery("limit", "5")
    limit, _ := strconv.Atoi(limitStr)
    age := c.Query("age")
    gender := c.Query("gender")
    country := c.Query("country")
    platform := c.Query("platform")
    var gen int
	if gender =="" {
		gen = 0
	}else if gender =="M"{
		gen = 1
	}else if gender =="F"{
		gen = 2
	}
    result := api.Select(offset,limit,age,gen,country,platform)
    c.JSON(http.StatusOK, result)
}
