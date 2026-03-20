FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY cmd/ cmd/
COPY apis/ apis/
COPY metrics/ metrics/

RUN go mod download

RUN CGO_ENABLED=0 go build -o fbdownloader .

FROM golang:1.24

WORKDIR /app

COPY --from=build /app/fbdownloader .

CMD ["/app/fbdownloader"]
