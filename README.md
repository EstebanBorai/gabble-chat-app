<div>
  <div align="center" style="display: block; text-align: center;">
    <img src="https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/go/go.png" height="120" width="120" />
  </div>
  <h1 align="center">gabble-chat-app</h1>
  <h4 align="center">💬 Tiny chat implementation made with Go and ReactJS</h4>
</div>

## Running Locally

### Using Go and Dep for local development

1. Make sure you already have *Go* and *dep* installed in your machine,
otherwise refer to the following sources to get them:
  - [Go](https://golang.org/dl/)
  - [dep](https://github.com/golang/dep)

2. Step into project root directory and run `dep ensure` to install
required packages:

```sh
cd ./gabble/server/
dep ensure
```

3. Run the application using `go run` command:

``` sh
go run ./src/main.go
```

### Running using Docker

1. Make sure you already have *Docker* installed in your machine,
otherwise install *Docker*:
  - [Docker](https://docs.docker.com/)

2. Build the **Dockerfile** available in the project's root directory:

```sh
cd ./gabble/server
docker build -t gabble-backend .
```

3. Finally run the container you just build using `docker run`:

```sh
docker run gabble-backend
```

### Contributions

Every contribution to this project is welcome.

### License

Licensed under the MIT License
