FROM golang:1.13.1

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN go build -o chi ./chi.go

CMD ["./chi"]