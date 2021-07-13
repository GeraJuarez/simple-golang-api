# Test stage TODO


# Build stage
FROM golang:1.16 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o kvs

# Copy binary
FROM scratch

COPY kvs .

# COPY *.pem .s

EXPOSE 8080

CMD ["/kvs"]
