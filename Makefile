# Имя бинарника
BINARY_NAME=myapp

# Точка входа
MAIN_FILE=./cmd/app/main.go

# Папка для бинарников
BIN_DIR=bin

#Папка со спецификацией
SPEC_DIR=./openapi/openapi.yaml

OAPI_CODEGEN_DIR=./openapi/oapi-codegen.yaml

.PHONY: all build run test clean help

# Цель по умолчанию
all: clean build

codegen:
	@oapi-codegen --config $(OAPI_CODEGEN_DIR) $(SPEC_DIR)
# Сборка
build: codegen
	@echo "Building $(BINARY_NAME)..."
	@if not exist $(BIN_DIR) mkdir $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME).exe $(MAIN_FILE)

run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BIN_DIR)\$(BINARY_NAME).exe

# Очистка
clean:
	@echo "Cleaning..."
	@rm -rf $(BIN_DIR)

# Справка
help:
	@echo "Available commands:"
	@echo "  make         - Build application"
	@echo "  make build   - Build binary"
	@echo "  make run     - Build and run"
	@echo "  make test    - Run tests"
	@echo "  make clean   - Remove binaries"