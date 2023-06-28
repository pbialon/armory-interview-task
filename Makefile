compile:
	go build -o build/logs_printer ./src/

run:
	go run ./src/logs_printer.go

test:
	go test ./...