FROM golang:alpine AS builder

# install git and update linux
RUN apk update && apk add git

# copy to container path and go to that path
COPY . $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/src/mypackage/myapp/

#get dependancies
#you can also use dep
RUN go get -d -v

#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/backoffice

FROM scratch

COPY --from=builder /go/bin/backoffice /go/bin/hello
ENTRYPOINT ["/go/bin/hello"]