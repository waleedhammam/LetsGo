# LetsGo
playground of distrubed go, a commit log service

## Run

`go run cmd/server/main.go`

## Examples

```bash
➜  server git:(main) curl -XPOST localhost:8080 -d '{"record": {"value": "c29tZXRlc3R2YWx1ZXM6MQ=="}}'
{"offset":0}
➜  server git:(main) curl -XPOST localhost:8080 -d '{"record": {"value": "c29tZXRlc3R2YWx1ZXM6Mg=="}}'
{"offset":1}
➜  server git:(main) curl -XPOST localhost:8080 -d '{"record": {"value": "c29tZXRlc3R2YWx1ZXM6Mw=="}}'
{"offset":2}
➜  server git:(main) curl -XGET localhost:8080 -d '{"offset":0}'
{"record":{"value":"c29tZXRlc3R2YWx1ZXM6MQ==","offset":0}}
➜  server git:(main) curl -XGET localhost:8080 -d '{"offset":1}'
{"record":{"value":"c29tZXRlc3R2YWx1ZXM6Mg==","offset":1}}
➜  server git:(main) curl -XGET localhost:8080 -d '{"offset":2}'
{"record":{"value":"c29tZXRlc3R2YWx1ZXM6Mw==","offset":2}}
```
