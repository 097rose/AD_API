package api

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"Api/model"
	"time"
)

func Create(data model.PostBody){
	
	db, err := sql.Open(
		"mysql",	
		"user01:000000@tcp(127.0.0.1:3306)/dcard",
	)
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping() 
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
	
	var gender int
	if data.Conditions[0].Gender =="" {
		gender = 0
	}else if data.Conditions[0].Gender =="M"{
		gender = 1
	}else if data.Conditions[0].Gender =="F"{
		gender = 2
	}
	startAtTime, _ := time.Parse(time.RFC3339, data.StartAt)

	endAtTime, _ := time.Parse(time.RFC3339, data.EndAt)
	st := startAtTime.Format("2006-01-02 15:04:05")
	et := endAtTime.Format("2006-01-02 15:04:05")

	result, err := db.Exec("INSERT INTO Ad (title, startAt, endAt, ageStart, ageEnd, gender) VALUES (?, ?, ?, ?, ?, ?)", data.Title, st, et, data.Conditions[0].AgeStart, data.Conditions[0].AgeEnd,gender)
	if err != nil {
		log.Fatal(err)
	}

	adID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	countries := data.Conditions[0].Country
	for _, country := range countries {
		_, err = db.Exec("INSERT IGNORE INTO Country (name) VALUES (?)", country)
		if err != nil {
			log.Fatal(err)
		}
	}

	platforms := data.Conditions[0].Platform
	for _, platform := range platforms {
		_, err = db.Exec("INSERT IGNORE INTO Platform (name) VALUES (?)", platform)
		if err != nil {
			log.Fatal(err)
		}
	}


	for _, country := range countries {
		_, err = db.Exec("INSERT INTO Ad_Country (adId, countryId) VALUES (?, (SELECT id FROM Country WHERE name = ?))", adID, country)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, platform := range platforms {
		_, err = db.Exec("INSERT INTO Ad_Platform (adId, platformId) VALUES (?, (SELECT id FROM Platform WHERE name = ?))", adID, platform)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Insert data success")

	

}