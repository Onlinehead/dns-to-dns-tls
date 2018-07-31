FROM onlinehead/golang-git-dep:1.9-alpine as build_env

COPY . /go/src/dns-to-dns-tls
RUN cd /go/src/dns-to-dns-tls \
    && go build

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=build_env /go/src/dns-to-dns-tls/dns-to-dns-tls /app/
COPY --from=build_env /go/src/dns-to-dns-tls/config.yaml /app/

EXPOSE 53
EXPOSE 53/udp
WORKDIR /app
ENTRYPOINT ["/app/dns-to-dns-tls"]




