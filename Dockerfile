FROM golang:1.21-alpine

WORKDIR /app

RUN export GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/

RUN go build -o /

EXPOSE 3000

CMD ["/salty"]
