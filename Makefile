build:
	@echo "Building..."
	@go build -o  bin/main cmd/main/main.go
run:build
	@echo "Running..."
	@./bin/main
