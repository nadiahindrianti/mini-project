FROM golang:1.19

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build -o mini-project .

CMD ["/mini-project"]