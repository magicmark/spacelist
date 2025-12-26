.PHONY: build install clean

build:
	go build -o spacelist

fmt:
	swift format -i *.swift
	go fmt .

clean:
	git clean -fdx

release-%:
	git tag -a "v$*"
	git push origin "v$*"

help:
	@echo "Available targets:"
	@echo "  build    - Build the spacelist binary"
	@echo "  install  - Build and install to /usr/local/bin/spacelist"
	@echo "  clean    - Remove built binaries"
