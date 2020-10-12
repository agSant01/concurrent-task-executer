DEFAULT_OUT_DIR = ./target
DEFAULT_OUT = main

OUT_DIR = $(DEFAULT_OUT_DIR)
OUT = $(DEFAULT_OUT)

.PHONY: build run


build:
	@ go build -v -o $(OUT_DIR)/ ./cmd/concurrent/main.go

run: build
	@ $(OUT_DIR)/$(OUT)