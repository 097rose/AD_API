package main

import (
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	// "log"
	// "net/http"
	// "Api/api"
	"Api/router"

)
func main() {
	r := router.SetupRouter()
	r.Run(":8800")

}