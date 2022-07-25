# Library Application

## How to run

1. Run PostgreSQL

```
$ docker run --name my-postgres -p 5432:5432 \
    -e POSTGRES_PASSWORD=1234 \
    -e POSTGRES_DB=library \
    -d postgres
```

2. Run Redis
```
$ docker run --name my-redis -p 6379:6379 -d redis
```

3. Run application
```
$ go run main.go
```

4. Open `frontend/index.html`
