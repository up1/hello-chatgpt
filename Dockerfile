FROM golang:1.19 as build
WORKDIR /app
COPY . /app
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure
RUN go build -o main .


FROM alpine:latest
COPY --from=build /app/main /app/main
EXPOSE 8080
CMD ["./main"]
