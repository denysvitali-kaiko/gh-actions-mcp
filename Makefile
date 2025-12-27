.PHONY: build test clean install run

BINARY=gh-actions-mcp

build:
	go build -o $(BINARY) .

test:
	go test ./... -v

clean:
	rm -f $(BINARY)

install:
	go install .

run:
	go run . --token=$$GITHUB_TOKEN

# For Claude Desktop MCP integration
install-mcp:
	cp $(BINARY) ~/.config/Claude\ Desktop/mcp-servers/

.PHONY: all
