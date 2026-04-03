# GoXpress
GoXpress provides an intuitive architecture for Go server projects, designed to facilitate rapid development of your API applications. It is bundled with variety of framework and packages to streamline your development process. While this setup is comprehensive, you are not required to utilize all included packages. Feel free to use this repository as a reference to tailor your project according to your specific needs.

|Stack|Package|
|-|-|
|Go `v1.25.0`|Echo `v5.1.0`|
|Air `v1.64.5`|Gorm `v1.31.1`|
|Postgres|EchoSwagger `v2.0.1`|
|Docker||

## Pre-Build Functions
- [x] Easy development with Air hot-reload
- [x] Easy deployment with docker
- [x] Gorm with postgres connection
- [x] Initial user table and user model
- [x] Request params validation
- [x] Register/Login password encryption
- [x] Simple Jwt auth token and authentication middleware
- [x] Default docs page and auth & user routes
- [x] Swagger api documentation (not auto generated)
- [x] Simple docker containerize

## Getting started
### Installation
1. Ensure that [Golang 1.25.0](https://go.dev/doc/install) is installed on system.
2. Go into project folder. (Linux Example: /www)
```sh
cd wwww
go run github.com/xncs120/goxpress@main project_name
```
### Setting up enviroment configs
1. Modify .env as necessary to suit your configuration requirements. There are several important key:
2. APP_ENV accept "development" / "production".
3. JWT_SECRET_KEY accept string generated with 256BITS_HMAC_ALGO_HS256 format.
4. DB_URL accept string for database connection (postgres). For other database support please visit gorm documentation.
### Migration
1. Database migration is in /main.go > db.Migration() that is commented. Uncomment it to use Gorm auto migrate function.
2. Remember to add any model struct to /db/migration.go whenever creating a new table for auto migrate to work.
### Serve api
1. Go into generated project folder and run command below.
2. Access on api browser http://localhost:8080/ or http://localhost:8080/docs after success.
3. The docs page title wont change to project_name, feel free to admend it in /views/index.html.
```sh
cd project_name
air -c air.toml
```
### Deploy with docker
```sh
# check did docker get .env values
docker-compose --env-file .env config
# build docker contianer
docker compose --env-file .env up --build
```

## Project layout (Default)
```sh
project_name/
├── client/ (frontend js framework)
├── config/
│   ├── app.go
│   ├── config.go
│   └── database.go
├── db/
│   ├── db.go
│   └── migration.go (auto migration)
├── handlers/
│   ├── auth.go
│   ├── landing.go
│   └── user.go
├── internal/ (any important function)
│   ├── request/
│   │   └── validator.go
│   └── security/
│       ├── password.go
│       └── token.go
├── models/
│   └── user.go (gorm struct or other gobal used struct)
├── router/
│   ├── api.go
│   ├── router.go
│   └── web.go
├── server/
│   └── main.go (server trigger)
├── views/ (frontend html)
│   ├── docs.go (swagger documentation)
│   │   └── root.yaml
│   ├── efs.go
│   └── index.html (homepage)
├── .env (server config)
├── .gitignore
├── air.toml (air config)
├── docker-compose.yaml
├── Dockerfile
├── go.mod
└── go.sum
```

## Reference and external source
- [Air Documentation](https://github.com/air-verse/air)
- [Echo Documentation](https://echo.labstack.com/)
- [Gorm Documentation](https://gorm.io/docs/index.html)
- [Echo Swagger Documentation](https://github.com/swaggo/echo-swagger)