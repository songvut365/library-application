# Library Application

## How to run

1. Run PostgreSQL

```
$ docker run --name my-postgres -p 5432:5432 -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=library -d postgres
```

2. Run Redis
```
$ docker run --name my-redis -p 6379:6379 -d redis
```

3. Run application
```
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