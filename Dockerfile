FROM golang:1.19.4-alpine3.17 as build
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["/app/main"]
