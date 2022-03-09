FROM docker.io/golang:latest as build

WORKDIR /go/src
COPY . .

RUN CGO_ENABLED=0 go build -o healthz

FROM scratch

COPY --from=build /go/src/healthz /

ENTRYPOINT ["/healthz"]
