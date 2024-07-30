FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o geoserver .

EXPOSE 8080

CMD ["./geoserver"]
