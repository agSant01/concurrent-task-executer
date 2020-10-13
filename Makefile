DEFAULT_OUT_DIR = ./target
DEFAULT_OUT = main.out

OUT_DIR = $(DEFAULT_OUT_DIR)
OUT = $(DEFAULT_OUT)

MAIN = ./main.go

.PHONY: build run

build:
	@ go build -v -o $(OUT_DIR)/$(DEFAULT_OUT) $(MAIN)

run: build
	@ $(OUT_DIR)/$(OUT)

test: 
	go test -timeout 30s ./...