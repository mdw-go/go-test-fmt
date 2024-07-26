# go-test-fmt

A library and command-line tool to format output of `go test ./...`.

---

Example output:

```
github.com/mdwhatcott/go-collections on git main [!] 
$ make | go-test-fmt
go version go1.23rc1 darwin/arm64
go fmt ./...
go mod tidy
go test -cover -count=1 -timeout=1s -race ./...
-  github.com/mdwhatcott/go-collections/internal/should        [missing test files]
ok github.com/mdwhatcott/go-collections/list            100.0%               0.279s
ok github.com/mdwhatcott/go-collections/queue           100.0%               0.464s
ok github.com/mdwhatcott/go-collections/set             100.0%               0.641s
ok github.com/mdwhatcott/go-collections/stack           100.0%               0.836s
```