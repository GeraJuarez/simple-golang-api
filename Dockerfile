# Test stage
FROM golang:1.16 as test
COPY . /src
WORKDIR /src
RUN go mod download
RUN go test ./...

# Build stage
FROM golang:1.16 as build
COPY . /src
COPY --from=test /go/pkg/mod/ /go/pkg/mod/ 
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o kvs

# Copy binary to container
FROM scratch
COPY --from=build /src/kvs .
COPY --from=build /src/.env .
# COPY *.pem .s
EXPOSE 8080
CMD ["/kvs"]
