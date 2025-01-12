.PHONY: all
all: sql css
	go build -o bin/main cmd/fuelr/main.go

.PHONY: sql
sql:
	sqlc generate

.PHONY: css
css:
	pnpm css

.PHONY: dev
dev: sql
	air & pnpm dev

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	pnpm test
