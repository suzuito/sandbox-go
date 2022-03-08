
# Protocol buffer

```bash
go run .
```

## メモ

### 構造体の再生成

```golang
protoc \
--go_out=. --go_opt="module=github.com/suzuito/sandbox-go/cmd/003" \
./data.proto
```