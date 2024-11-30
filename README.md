# GoXpress
GoXpress provides an opinionated architecture for Go server projects, designed to facilitate rapid development of your API applications. It is bundled with variety of framework and packages to streamline your development process. While this setup is comprehensive, you are not required to utilize all included packages; feel free to use this repository as a reference to tailor your project according to your specific needs.
|Stack|Package|
|-|-|
|Go `v1.23.2`|Echo `v4.12.0`|
|Postgres|Goose `v3.22.1`|
|Docker|Gorm `v1.25.12`|
||Air (Go Hot-Reload)|

## Pre-Build Functions
[x] Makefile (Docker command, Migration command, Go command)
[x] Easy development/production enviroment setup (Env configs, Docker configs, Air configs)
[x] Go embeded assets (statics, templates)
[x] Home page, Register user page, Sample api schema page
[x] Api routes for user
[x] Initial User table
[x] Simple Jwt authentication

## Getting started
### Installation
1. Ensure that [Golang 1.23](https://go.dev/doc/install) and [Docker](https://www.docker.com/products/docker-desktop/) is installed on system.
2. Go into project folder. (Linux Example: /www)
```sh
cd wwww
go run github.com/xncs120/goxpress@master project-name
```
### Setting up enviroment configs
1. Modify .env as necessary to suit your configuration requirements. There are several important key:
-- APP_ENV accept "development" / "production".
-- JWT_SECRET accept string generated with 256BITS_HMAC_ALGO_HS256 format.
2. Examine the air.toml and Docker configuration files to determine if any adjustments are needed for your specific setup.
3. Launch the database and application service container (P.S: Start database first then app)
```sh
make docker-db-up
make docker-app-up
```
### Migration
1. Access into docker terminal and do users table migration using goose. The reason migration need to be run manually is because there might be time that some table is too large to alter and will take really long time to migrate which cause deployment of app slow. (P.S: Example of users table shown is pre-builded)
```sh
make goose-create add_users_table

// go to example: ./database/migrations/20241015105036_add_users_table.sql

-- +goose Up
CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	username TEXT UNIQUE NOT NULL,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	status INTEGER NOT NULL DEFAULT 1,
	created_at TIMESTAMP NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;

make goose-up
```

## Project layout
```sh
project/
├── assets/
│   ├── statics/
│   │   └── styles.css
│   ├── templates/
│   │   └── index.html
│   └── efs.go
├── cmd/
│   └── api/
│       └── main.go
├── goose/
│   ├── migrations/
│   └── seeders/
├── internal/
│   ├── base/
│   │   ├── auth/
│   │   │   └── auth.go
│   │   ├── config/
│   │   │   ├── config.go
│   │   │   └── list.go
│   │   ├── database/
│   │   │   └── database.go
│   │   ├── resource/
│   │   │   └── resource.go
│   │   └── routes/
│   │       └── routes.go
│   ├── landing/
│   │   └── handlers.go
│   └── user/
│       ├── handlers.go
│       └── model.go
├── .env
├── .gitignore
├── air.toml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── Makefile
```

## Reference and external source
- [Organizing a Go module](https://go.dev/doc/modules/layout)
- [Golang Documentation](https://go.dev/doc/effective_go)
- [Echo Documentation](https://echo.labstack.com/)
- [Goose Documentation](https://pressly.github.io/goose/)
- [Gorm Documentation](https://gorm.io/docs/index.html)
- [Air Documentation](https://github.com/air-verse/air)