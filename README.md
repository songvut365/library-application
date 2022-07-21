# Library Application

## How to start

```
$ docker run --name my-postgres -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=library -d postgres
$ docker run --name my-redis -d redis
$ go run main.go
```

## Directory Structure

```
library-application
├── README.md
├── config
│   ├── database.go
│   └── redis.go
├── frontend
│   ├── book.html
│   ├── book_detail.html
│   ├── borrow_list.html
│   ├── index.html
│   ├── member.html
│   └── member_list.html
├── go.mod
├── go.sum
├── handler
│   ├── book.go
│   ├── book_detail.go
│   ├── borrow.go
│   ├── member.go
│   └── state.go
├── main.go
└── model
    └── models.go
```