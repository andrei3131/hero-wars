
FROM golang:latest

LABEL maintainer="Andrei-Octavian Brabete <brabete.andrei96@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app
RUN make build
RUN make test

# Remove the boot command if you want an interactive console to your container.
CMD ["./bin/hero-wars"]
