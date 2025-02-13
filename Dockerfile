FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.15
WORKDIR /app

COPY --from=build /app/main /app/

EXPOSE 3001

CMD ["./main"]
