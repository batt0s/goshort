# Run temp
run:
	go run .

# test all
test:
	go test -v ./tests/database_test.go
	go test -v ./tests/shortener_test.go