FROM golang:alpine
WORKDIR /build

COPY . .
RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -o workshop-product-service -ldflags "-X main.version=`cat version` -X main.build=`date -u '+%Y-%m-%d-%H%M%S'`" ./cmd/main.go

FROM golang:alpine
WORKDIR /app

COPY --from=0 /build/workshop-product-service .

CMD ["/app/workshop-product-service"]