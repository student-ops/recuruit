# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR /app
COPY ./content/ ./
RUN go mod download
RUN go build -o /app/main
EXPOSE 9000
CMD ["/app/main"]
