## Fakerelic

This project is intended to monitor web applications that are in production

### Setup

* Docker
```bash
docker-compose build
```
* MacOS

Download Golang from [Website](https://golang.org/dl/) and extract files to `/usr/local/go`

* Linux

Download Golang Binaries from [Website](https://golang.org/dl/) and extract files to `/usr/local/go/bin`. Add executable to PATH:

```bash
export PATH=$PATH:/usr/local/go/bin
```

### Run

With Docker

```bash
docker-compose up app
```

Native OS

```bash
go get ./...
go run cmd/main.go
```

### Run Tests

With Docker

```bash
docker-compose up test
```

Native OS

```bash
go test ./...
```

### Run Lint

```bash
golangci-lint run -v
```
