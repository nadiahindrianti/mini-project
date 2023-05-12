FROM golang:1.19

RUN mkdir /app

WORKDIR /app

COPY go.mod /app

COPY go.sum /app

RUN go mod download

ADD . /app

COPY *.go ./

EXPOSE 8080

RUN go build -o main .

CMD ["/app/main"]