.PHONY: build
build: 
	@go build
	@echo "cmdr built"

.PHONY: install
install: build
	@sudo install -d /usr/local/bin/ && \
	sudo install ./cmdr /usr/local/bin/ && \
	rm cmdr
	@echo "cmdr installed"
