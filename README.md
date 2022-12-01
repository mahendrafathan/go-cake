# go-cake
Privy Take Home Test.
This project `go-cake` is api for cake
## Owner Info
Mahendra Fakhrul Fathan (mahendrafathan@gmail.com)

## Prerequisite
Please install [Docker](docker.com/get-started) then install [Docker Compose](https://docs.docker.com/compose/install/) on your laptop to run this app easily.

## Getting Started
### How To Run
1. Unzip into folder
2. Change directory to the folder `cd /go-cake`
3. Run the app using docker compose: `docker-compose up --build -d`, this might takes 2 - 3 mins, so please wait
5. You're good to go! The app will serve on port `9090`, please kindly check api curl documentation at the end of this page
6. You can check db using admirer, it's serve on `8080`

### Docker Compose Explanation
1. Docker will build and set Mysql DB Based on `.env` file
2. Docker will build and run golang application
⋅⋅* Download all application dependencies
⋅⋅* Run unit test from `unitTest.sh`, if unit test failed or coverage below 80%, application will not running

### Run Unit Test Separately
If you want to see how unit test works and coverage please run `chmod +x ./unitTest.sh && ./unitTest.sh`

### Customize DB 
Please change on `.env` file

## App Info
### Directory App
```bash
.
├── Dockerfile
├── README.md
├── common
│   ├── db
│   │   └── database.go
│   └── logger
│       └── logger.go
├── controllers
│   ├── cake.go
│   └── interfaces.go
├── doc
│   └── openapi.yaml
├── docker-compose.yml
├── go.mod
├── go.sum
├── helpers
│   └── response.go
├── init
│   └── init.sql
├── main.go
├── makefile
├── models
│   └── cake.go
├── repositories
│   ├── cake.go
│   └── cake_mysql.go
├── routes
│   ├── middleware.go
│   └── router.go
└── unitTest.sh
```