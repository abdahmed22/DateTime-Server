FROM golang:1.22.5-alpine3.20 AS base

WORKDIR /app

COPY go.mod go.sum ./

COPY pkg/httpserver ./pkg/httpserver/
COPY cmd/httpserver/main.go ./cmd/httpserver/

COPY pkg/ginserver ./pkg/ginserver/
COPY cmd/ginserver/main.go ./cmd/ginserver/

ENV PORT=8080

RUN go mod download

FROM base AS build-http
RUN go build -o /bin/httpserver ./cmd/httpserver


FROM base AS build-gin
RUN go build -o /bin/ginserver ./cmd/ginserver


FROM scratch AS http
COPY --from=build-http /bin/httpserver /cmd/

ENTRYPOINT [ "/bin/httpserver" ]

EXPOSE 8080

CMD ["./main"]

FROM scratch AS gin
COPY --from=build-gin /bin/ginserver /cmd/

ENTRYPOINT [ "/bin/ginserver" ]

EXPOSE 8080

CMD ["./main"]
