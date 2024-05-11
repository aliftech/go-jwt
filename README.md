# GO-JWT

This project is created as an exercise for unserstanding the implementation of `Json Web Token (JWT)` in golang signin rest api. This project is the continuation of the previous project (https://github.com/aliftech/todo-api)

## Instalation

### Pre-requisites

Before installing this project, you should know about the project requirements. Here are the requirements that you need to prepared:

1. Go Programming Laguage
2. Knowledge about Go.
3. Knowledge about REST-API

### Dependencies

This project is build using some third parties or dependencies. Here are the required dependencies:

1. Gin (https://gin-gonic.com/)
2. Gorm (https://gorm.io/)
3. Go dotenv (https://github.com/joho/godotenv)
4. CompileDaemon (https://github.com/githubnemo/CompileDaemon)
5. MySQL driver (https://gorm.io/docs/connecting_to_the_database.html)
6. Go-JWT (https://pkg.go.dev/github.com/golang-jwt/jwt/v5)
7. Go Cryptography (https://pkg.go.dev/golang.org/x/crypto). in this case we use bcrypt to encrypt our password.

### Database Migration

We make a little modification, so that you don't need to run the migration manually. When you running this project, it will automaticly run the migration.

### Running The Project

To run the project, you can use a usual go command:

```bash
go run main.go
```
