## KV
KV is a toy in-memory key value store built primarily in an effort to write more go and check out `grpc`.
This is still a work in progress.

```
// download dependencies
go mod download
// to run the server (listens on port 20020 by default)
go run server/main.go -port 12345
// to start the CLI
go run cli/cli.go -port 12345
```
The CLI consumes the client library, and currently supports the following commands:
```
insert - Insert key and value into kv
lookup - Lookup key in kv
delete - Delete key in kv
help   - Show available commands
clear  - Clear the terminal screen
exit   - Close your kv client
```
