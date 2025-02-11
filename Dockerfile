FROM golang:alpine
WORKDIR /app
ENV TELEGRAM_API_TOKEN=YOUR_API_TOKEN
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./main.go
EXPOSE 8080
CMD ["./main"]
