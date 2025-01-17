FROM golang:1.23.4-alpine3.20

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o todo .

EXPOSE 8000

CMD [ "./todo" ]