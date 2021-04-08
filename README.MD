# Go Project Structure #

**This is a sample structure for managing Go projects which provides versioning for REST APIs. This repository has following packages which makes it useful while creating your first GO project.**
- gorm : It is the ORM library in Go which provides user friendly functions to interact with database. It supports features like ORM, Associations, Hooks, Preloading, Transaction, Auto Migration, Logger etc.
- gin : Gin is a web framework for Go language. Here gin is used for increase performance and productivity.
- godotenv : Basically used for load env variables from .env file.
- mysql : It provides the mysql driver for connecting Go with MySQL.


## Structure ##
<img src="https://raw.githubusercontent.com/Mindinventory/Golang-Project-Structure/master/structure.png" width="400" style="max-width:100%;">

## Directories ##

- ApiHelpers
  - Basically contains the helper functions used in returning api responses, HTTP status codes, default messages etc.
  
- Controllers
  - Contains handler functions for particular route to be called when an api is called.

- Helpers
  - Contains helper functions used in all apis

- Middlewares
  - Middleware to be used for the project

- Models
 - Database tables to be used as models struct

- Resources
  - Resources contains all structures other than models which can be used as responses

- Routers
  - Resources define the routes for your project

- Seeder
  - It is optional, but if you want to insert lots of dummy records in your database, then you can use seeder.

- Services
  - All the core apis for your projects should be within services.

- Storage
  - It is generally for storage purpose.

- Templates
  - Contains the HTML templates used in your project

- env

## RUN THE SERVER (Basic commands) ##
```go
. For running the server you have to run following command go run main.go It will start your server at the port you have mentioned in .env file.

. If you want to run the server in port other than default, then run following command go run main.go <specific port>

. If you want to create a build for your project and upload in server, then you have to run following command go build
```

## API Version ##

```go
127.0.0.1:8099/api/v1/user-list

127.0.0.1:8099/api/v2/user-list


```
