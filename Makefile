.PHONY install

install: ## builds and copies the binary to your bin directory
	@go build -o beamer cmd/main.go
	@cp ./beamer /usr/local/bin
