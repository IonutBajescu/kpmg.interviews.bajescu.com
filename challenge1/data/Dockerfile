FROM golang:1.14
WORKDIR /go/src/app
COPY . .
RUN go build -v -ldflags "-linkmode external -extldflags -static" .

FROM alpine
COPY --from=0 /go/src/app/app /bin/app
CMD ["/bin/app"]

