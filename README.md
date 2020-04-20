# gabble-backend

## Running Locally

### Using Go and Dep for local development

1. Make sure you already have *Go* and *dep* installed in your machine,
otherwise refer to the following sources to get them:
  - [Go](https://golang.org/dl/)
  - [dep](https://github.com/golang/dep)

2. Step into project root directory and run `dep ensure` to install
required packages:

```sh
cd ./gabble-backend/
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
cd ./gabble-backend/
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
