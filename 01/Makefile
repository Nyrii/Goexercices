BINARY_NAME=01
TARGET_DIR=target

all: test build

## Test:
test: ## Runs the tests of the project
	go test ./...

test_coverage: ## Runs the tests of the project and exports the coverage.
	mkdir -p ${TARGET_DIR}
	go test ./... -coverprofile=${TARGET_DIR}/coverage.out
	@echo "Exported the test coverage outcome file in directory='${TARGET_DIR}'."

## Build:
build: ## Builds the project and put the output binary in target/bin/
	mkdir -p ${TARGET_DIR}
	GO111MODULE=on go build -o ${TARGET_DIR}/$(BINARY_NAME)
	@echo "Build finished; binary='${BINARY_NAME}' available in directory='${TARGET_DIR}'."

clean: ## Removes the build related files and directories
	rm -rf ./target
