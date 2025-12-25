.PHONY: build install clean

build:
	go build -o sl

install: build
	cp sl /usr/local/bin/sl

fmt:
	swift format -i *.swift

clean:
	rm -f sl

help:
	@echo "Available targets:"
	@echo "  build    - Build the spacelist binary as 'sl'"
	@echo "  install  - Build and install to /usr/local/bin/sl"
	@echo "  clean    - Remove built binaries"
