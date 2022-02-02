PROJECT = github.com/readygo67/LiquidationBot
TARGET_DIR = bin

.PHONY: all build venusd venuscli

all: build

build: venusd venuscli

venusd:
	go build -o $(TARGET_DIR)/venusd $(PROJECT)/cmd/venusd

venuscli:
	go build -o $(TARGET_DIR)/venuscli $(PROJECT)/cmd/venuscli