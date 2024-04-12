# Variables
APP_NAME := clicker
BIN_DIR := bin
CMD_DIR := cmd
MAIN_FILE := $(CMD_DIR)/main.go

# Commands
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "Build complete. Binary is located at $(BIN_DIR)/$(APP_NAME)"

delete:
	@echo "Deleting $(APP_NAME) binary..."
	@rm -f $(BIN_DIR)/$(APP_NAME)
	@echo "$(APP_NAME) binary deleted."

run:
	@echo "Starting $(APP_NAME)..."
	@$(BIN_DIR)/$(APP_NAME)

rebuild: delete build

.PHONY: build delete rebuild
