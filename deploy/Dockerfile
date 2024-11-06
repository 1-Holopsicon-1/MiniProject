FROM golang:1.23.2-bullseye

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

EXPOSE 5000

CMD ["go", "run", "cmd/main.go", "-migrate", "-rnddata", "-start"]