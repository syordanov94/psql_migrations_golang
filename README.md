# SQL Migrations in Go

A small lab that performs database migrations to a local PSQL database innitialized using Docker. To perform the migrations, we use the [go-migrate](https://github.com/golang-migrate/migrate) library. 

## Prerequisites

- Golang 1.19 or higher installed 
- Docker installed
- _Recomended but not mandatory_ VS Code or a similiar IDE 

## How to install and Run the project

- First you will have to clone the project from this github repository.

```bash
git clone https://github.com/syordanov94/go_redis_concurrency.git
```

- Once cloned, access the root of the project and innitiate the local **PSQL** instance using **Docker**. This can be done buy running the following command:

```bash
docker compose up
```

You can add (or modify) .**.sql** files in the *migrations* folder if you want to test custom migrations.

Once the database is innitialized, you can run the project and apply the migrations to your local database by running:

```bash
go run cmd/main.go
```

## How to test the project
