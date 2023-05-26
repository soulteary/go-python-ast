# Go Python AST

Get Python AST with Go.

## Usage

Start Go Python AST service, `8080` for HTTP client and `8081` for GRPC client.

```bash
# docker image base debian
docker run --rm -it -p 8080:8080 -p 8081:8081 soulteary/go-python-ast
# or alpine base
docker run --rm -it -p 8080:8080 -p 8081:8081 soulteary/go-python-ast:alpine
```

You can use the following commands to test the HTTP API.

```bash
# curl --request POST 'http://127.0.0.1:8080/api/convert' --header 'Content-Type: application/json' --data-raw 'print("hello world")'

"{\"Module\": {\"body\": [{\"Expr\": {\"value\": {\"Call\": {\"func\": {\"Name\": {\"id\": \"print\", \"ctx\": \"Load\"}}, \"args\": [{\"Str\": {\"s\": \"hello world\"}}], \"keywords\": []}}}}]}}"
```

You can use the following commands to test the GRPC API.

```bash
cd example/grpc-client
go run main.go
```

## Tutorial

- [Using Golang and Docker to implement Python computing services](https://soulteary.com/2023/05/22/using-golang-and-docker-to-implement-python-computing-services.html)

## Dev

```bash
docker build -t soulteary/go-python-ast . -f docker/Dockerfile
```

## Credits

- [Docker Python in Go](https://github.com/soulteary/docker-python-in-go), the principle.
