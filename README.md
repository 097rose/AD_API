# AD_API
2024 Dcard Backend Intern Assignment

## Description
Implement two API services using Golang and MySQL as the database.  
Including two API services, one for POST requests to generate advertisements and another for GET requests to retrieve desired advertisements.

## Requirement
### go 1.21.6
Using go.mod files to manage dependencies.
```
go mod tidy
```
 
### MySQL 8.3
Log in to the MySQL server and enter the password.
```
mysql -u root -p
```
Initialize the database and create a user.
```
source YourPath/ad.sql
```
## Folder Structure
```
AD_API/
|-- api/
|   |-- create_data_test.go
|   |-- create_data.go
|   |-- select_data_test.go
|   |-- select_data.go
|-- model/
|   |-- model.go
|-- router/
|   |-- router_test.go
|   |-- router.go
|-- ad.sql
|-- ad._test.postman.collection
|-- data.json
|-- go.mod
|-- go.sum
|-- main.go
|-- README.md (thisfile)
```
- api：Define the complete function.  
- model：Define required  structures.  
- router：Set up the routing configuration for the API endpoints.  

## Run
Run main.go to build API server
- port:8800

```
go run main.go
```
POST request - use data.json for request body
```
curl -X POST -H "Content-Type: application/json" -d "@data.json" http://localhost:8800/api/v1/ad

```
GET request 
- parameter : offset, limit, age, gender, country, platform
```
curl http://127.0.0.1:8800/api/v1/ad?limit=8&age=20

```
### Postman collection
- create_ad : for post request
- select_ad : for get request

## Unit test
- TestCreste
- TestSelect
- TestSetupRouter_GET
- TestSetupRouter_PostRequest
