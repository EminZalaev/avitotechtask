FROM golang:1.17-alpine
WORKDIR /avitotechtask
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./app ./cmd/app
EXPOSE 8080
CMD ["/avitotechtask/app"]
