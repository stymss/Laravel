BINARY_NAME=larvyApp

build:
	@go mod vendor
	@echo "Building Larvy..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Larvy built!"

run: build
	@echo "Starting Larvy..."
	@./tmp/${BINARY_NAME} &
	@echo "Larvy started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Larvy..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Larvy!"

restart: stop start