FROM golang:alpine AS go-builder

RUN apk update && apk add --no-cache git

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/github.com/whizzes/gabble-backend

COPY . .

RUN dep ensure

RUN ls

RUN ls ./src

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/gabble-backend ./src/main.go

FROM scratch

COPY --from=go-builder /go/bin/gabble-backend /go/bin/gabble-backend

ENTRYPOINT ["/go/bin/gabble-backend"]
