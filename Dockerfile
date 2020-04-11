FROM golang:1

WORKDIR /app

# COPY . .

CMD ["go", "run", "main.go"]
