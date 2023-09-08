***This project is my example of Golang API server.***

## How to run

1. Install golang.
2. Install docker.
3. Install docker-compose.
4. Run `docker-compose up -d` to run in docker.

## How to test

1. Run `go test ./...` to run all tests.
2. Run `go test ./... -cover` to run all tests with coverage.
3. Run `go test ./... -coverprofile=coverage.out` to run all tests with coverage and save it to file.
4. Run `go tool cover -html=coverage.out` to open coverage in browser.
5. Run `go test ./... -bench=. -benchmem` to run all benchmarks.
6. Run `go test ./... -bench=. -benchmem -cpuprofile=cpu.out` to run all benchmarks with cpu profile.
7. Run `go tool pprof -http=:8080 cpu.out` to open cpu profile in browser.
8. Run `go test ./... -bench=. -benchmem -memprofile=mem.out` to run all benchmarks with memory profile.
9. Run `go tool pprof -http=:8080 mem.out` to open memory profile in browser.
10. Run `go test ./... -bench=. -benchmem -memprofile=mem.out -cpuprofile=cpu.out` to run all benchmarks with memory and cpu profiles.
