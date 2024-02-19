build:
	go build -o ./bin/goshort .

run:
	./bin/goshort

# Run temp
dev:
	go run .

# test all
test:
	go test -v ./tests/database_test.go
	go test -v ./tests/controller_test.go