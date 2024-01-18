# On MACOS requires gnu-sed. Run `brew info gnu-sed` and follow instructions to replace default sed.
# Negative lookbehind tries to find "= `" pattern to not affect go templates for code generation
imports:
	find . -type f -name '*.go' -exec sed -zi 's/(?<== `\s+)"\n\+\t"/"\n"/g' {} +
	goimports -local "github.com/aiven/aiven-go-client-v3" -w .

GEN_HANDLER_DIR ?= handler

go-generate:
	rm -rf $(GEN_HANDLER_DIR)
	GEN_HANDLER_DIR=$(GEN_HANDLER_DIR) go run -tags=generator ./generator/...

generate: go-generate imports

test:
	go test -v ./...
