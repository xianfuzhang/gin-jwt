FROM golang:1.18 AS build-stage

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace

COPY . .

RUN go mod download \
    && make build
 
FROM alpine:3.14 AS release-stage

# ENV GIN_MODE=release

WORKDIR /workspace

COPY /scripts/sqlite.sql .

COPY --from=build-stage /workspace/_output/clean_app .

RUN apk add sqlite \
    && cat sqlite.sql | /usr/bin/sqlite3 sqlite.db

EXPOSE 8080

CMD ["./clean_app"]

