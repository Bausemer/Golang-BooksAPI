# About this project

This project is based on the book <u>*"Let's Go" by Alex Edwards*</u> and serves the purpose of learning the programming language Golang.

## Useful commands

### 1. Write logs into log files at runtime

```bash
go run ./cmd/web >>./tmp/info.log 2>>./tmp/error.log
```


go.mod
go.sum
Dockerfile
cmd/web/handlers.go
cmd/web/helpers.go
cmd/web/main.go
cmd/web/routes.go
internal/models/error.go
internal/models/snippets.go
tmp/
ui/html/...
ui/static/...