# base stage that copy files and builds them

FROM golang:1.22.5-alpine3.20 AS base

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /bin/httpserver ./cmd/httpserver/main.go
RUN go build -o /bin/ginserver ./cmd/ginserver/main.go

# run stage that run binaries
FROM scratch AS http
COPY --from=base /bin/httpserver /bin/
ENTRYPOINT [ "/bin/httpserver" ]

EXPOSE 8090

FROM scratch AS gin
COPY --from=base /bin/ginserver /bin/
ENTRYPOINT [ "/bin/ginserver" ]

EXPOSE 8085

