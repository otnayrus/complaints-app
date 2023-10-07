FROM golang:1.21-alpine as Build
COPY . .
RUN GOPATH= go build -o /main app/main.go
FROM alpine:latest
COPY --from=Build /main .
EXPOSE 8080
ENTRYPOINT ["./main"]