FROM golang:1.11.6 as builder

ARG VERSION
ENV CGO_ENABLED=0

WORKDIR $GOPATH/src/md5go

COPY . .
 
RUN go install -a --installsuffix cgo -ldflags "-X main.Version=${VERSION}" ./...

FROM iron/base

COPY --from=builder /go/bin/md5go /usr/local/bin
RUN chmod a+x /usr/local/bin/md5go

ENTRYPOINT ["/usr/local/bin/md5go"]