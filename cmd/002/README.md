# cel-go example

```bash
air
```

```golang
protoc \
--go_out=. --go_opt="module=github.com/suzuito/sandbox-go/cmd/002" \
--go-grpc_out=. --go-grpc_opt="module=github.com/suzuito/sandbox-go/cmd/002" \
./data.proto
```
