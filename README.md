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
|-- data.json
|-- go.mod
|-- go.sum
|-- main.go
|-- README.md (thisfile)
```
api：Define the complete function.
model：Define required  structures.
router：Set up the routing configuration for the API endpoints.

## Run
```
  $  cd LAIAIAI
```
```
  $  python main.py
```
```
  $  [enter your image path]
```
