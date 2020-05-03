FROM golang:alpine AS go-builder

RUN apk update && apk add --no-cache git

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/github.com/whizzes/gabble/server

COPY . .

RUN dep ensure

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/gabble/server ./main.go

RUN /go/bin/gabble/server

FROM scratch

COPY --from=go-builder /go/bin/gabble/server .

ENTRYPOINT ["/go/bin/gabble/server/main"]
