.PHONY: all
all: css
	go build -o bin/main

.PHONY: css
css:
	pnpm css

.PHONY: dev
dev:
	air & pnpm dev

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	pnpm test
