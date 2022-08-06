FROM golang:1.18-alpine3.15

RUN apk add git
RUN apk add build-base

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /out/dist

EXPOSE 3000

CMD ["/out/dist"]