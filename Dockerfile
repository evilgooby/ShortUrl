FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .. .

RUN go build ./cmd/main.go

ENTRYPOINT ["go", "run", "./cmd", "main.go"]