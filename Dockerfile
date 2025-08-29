# Build stage
FROM golang:1.24-alpine AS builder

# work directory inside the image
WORKDIR /app

# copy all files from current work dir to /app
COPY . .

RUN go build -o main main.go

#Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
EXPOSE 3000

CMD ["/app/main"]