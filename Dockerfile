FROM golang:latest

WORKDIR /app
COPY pkg ./pkg
COPY cli ./cli
COPY go.mod ./
RUN go mod tidy -e

Run cd cli && go build -o pub

CMD ["/app/cli/pub"]