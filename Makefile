.PHONY: client

LOG = $(shell printf "\033[34;1m▶ \033[0m")
DIST_DIR = "./dist/"

npm: $(info $(LOG) Ensuring client dependencies are up to date...)
	cd ./client && npm install

client: $(info $(LOG) Running client locally)
	cd ./client && npm start

dep: $(info $(LOG) Ensuring server dependencies are up to date...)
	cd ./server && dep ensure -v

server: dep; $(info $(LOG) Running server locally)
	go run ./server/main.go

env: $(info $(LOG) Creating an .env file...)
	bash ./build-env.sh

build: dep ; $(info $(LOG) Golang is building in `gabble/server/dist`)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/server ./server/main.go
