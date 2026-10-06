FROM golang:1.27-alpine

WORKDIR /app

COPY . .

EXPOSE 3000

ENTRYPOINT ["go", "run", "main.go"]