FROM golang:alpine as swagger
LABEL stage=sawgbuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN go install github.com/swaggo/swag/cmd/swag@latest

FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .
COPY --from=swagger /go/bin/swag /go/bin/

RUN swag init --parseDependency --parseInternal -g main.go --output docs/
RUN go build -ldflags="-s -w" -o ./main ./main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/America/New_York /usr/share/zoneinfo/America/New_York

ENV TZ America/New_York

WORKDIR /app

COPY --from=builder /build/main /app/
COPY --from=builder /build/.env.local /app/
COPY --from=builder /build/docs/ /app/docs

EXPOSE 8000

CMD ["/app/main"]
