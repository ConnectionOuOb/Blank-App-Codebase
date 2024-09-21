FROM golang:1.18

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o backend .

EXPOSE 8080

CMD ["./backend"]
