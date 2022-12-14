include .env
export

# ==============================================================================
# Main
run: 
	go run cmd/app/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...