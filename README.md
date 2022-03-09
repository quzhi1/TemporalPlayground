# TemporalPlayground
This is my playground for temporal

Following https://docs.temporal.io/docs/go/hello-world-tutorial

## Unit testing
```bash
go test
```

## Manual run
```bash
go run worker1/main.go # Terminal will freeze
# Open another terminal
go run worker2/main.go # Terminal will freeze
# Open another terminal
go run start/main.go
```