FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN  go mod download

COPY . .

RUN go build -o app cmd/main.go
ENTRYPOINT [ "./app" ]
EXPOSE 3000
