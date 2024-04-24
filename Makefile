GOOS=linux
GOARCH=arm64
BUILD_DIR=bin

build:
	@echo "Compiling the code for processDocuments $(GOOS)/$(GOARCH)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go mod tidy
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/basic-lambda-startup/bootstrap cmd/main.go
	@echo "Compilation of processDocuments concluded"

clean:
	@echo "Limpando..."
	rm -rf $(BUILD_DIR)

zip: clean build
	@echo "Zipping Lambda functions"
	@mkdir -p $(BUILD_DIR)
	@zip -j $(BUILD_DIR)/basic-lambda-startup.zip $(BUILD_DIR)/basic-lambda-startup/bootstrap
	@echo "ZIP files created"