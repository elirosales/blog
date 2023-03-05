FROM golang:1.19.6-alpine

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /blog

EXPOSE 8080

CMD [ "/blog" ]
