FROM golang:1.17 as builder

WORKDIR /build
COPY . .
# go mod net will using c libary  , stop that by using build -tags netgo
RUN go build -tags netgo -o ReverseProxyServer
RUN chmod +x ReverseProxyServer

FROM scratch

EXPOSE 80
WORKDIR /server
COPY --from=builder /build/ReverseProxyServer .
COPY --from=builder /build/config.json .
CMD ["/server/ReverseProxyServer"]