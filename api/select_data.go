package api

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	// "net/http"
	// "strings"
	"strconv"
	"Api/model"
	// "github.com/gin-gonic/gin"
)

func Select(offset int, limit int, age string, gender int,country string, platform string) model.ResultBody{
	db, err := sql.Open(
		"mysql",	
		"user01:000000@tcp(127.0.0.1:3306)/dcard",
	)
	
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping() 
	if err != nil {
		log.Fatalln(err)
	}

	sqlQuery := "SELECT DISTINCT a.title, a.startAt, a.endAt FROM ad a LEFT JOIN ad_Country ac ON a.id = ac.adId LEFT JOIN Country co ON ac.countryId = co.id LEFT JOIN ad_Platform ap ON a.id = ap.adId LEFT JOIN Platform p ON ap.platformId = p.id"
	sqlQuery += " WHERE 1=1"
	if age != "" {
		ageint,_ := strconv.Atoi(age)
		sqlQuery += fmt.Sprintf(" AND %d BETWEEN a.ageStart AND a.ageEnd", ageint)
	}
	//select active ad
	sqlQuery += " AND NOW() BETWEEN a.startAt AND a.endAt "
	sqlQuery += fmt.Sprintf(" AND a.gender = %d", gender)
	//select country 
	if country != "" {
		sqlQuery += fmt.Sprintf(" AND co.name = '%s' OR co.name IS NULL", country)
	}
	//select platform
	if platform != "" {
		sqlQuery += fmt.Sprintf(" AND p.name = '%s' OR p.name IS NULL", platform)
	}
	//show by acs
	sqlQuery += " ORDER BY a.endAt ASC"
	//limit and offset
	sqlQuery += fmt.Sprintf(" LIMIT %d OFFSET %d;", limit,offset)
	fmt.Println(sqlQuery)

	rows, err := db.Query(sqlQuery)
	if err != nil {
		fmt.Println("Error querying database:", err)
	}
	var AllData model.ResultBody
	for rows.Next() {
		var Data model.ResultItem
		err := rows.Scan(&Data.Title, &Data.StartAt, &Data.EndAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		AllData.Items = append(AllData.Items, Data)
    }
	return AllData
	

}