FROM golang:1.16 AS builder

WORKDIR /go/src/getir-case-study
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/getir-case-study/app .
CMD ["./app"]