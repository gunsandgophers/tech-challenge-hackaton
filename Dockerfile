FROM golang:1.23-alpine3.20

ENV PROJECT_DIR=/api \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /api

# Download Go modules
COPY . .

RUN go mod download

RUN go build -o build/api cmd/api/main.go

CMD [ "build/api" ] 

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080
