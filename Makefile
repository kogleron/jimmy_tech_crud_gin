GO       ?= go
GOFLAGS  ?=

docker:
	docker compose up

format:
	@echo "Formatting..."
	${HOME}/go/bin/gofumpt -l -w -extra .
	@echo "Formatting imports..."
	@for f in $$(find . -name '*.go'); do \
		${HOME}/go/bin/goimports-reviser -file-path $$f -project-name jimmy_tech_crud_gin; \
	done

lint:
	@echo "Linting"
	${HOME}/go/bin/golangci-lint run

install:
	@echo "Install required programs"
	$(GO) $(GOFLAG) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) $(GOFLAG) install golang.org/x/tools/cmd/goimports@latest
	$(GO) $(GOFLAG) install mvdan.cc/gofumpt@latest
	$(GO) $(GOFLAG) install github.com/vektra/mockery/v2@latest
	$(GO) $(GOFLAG) get -v github.com/incu6us/goimports-reviser
	$(GO) $(GOFLAG) get -v github.com/swaggo/swag/cmd/swag
	$(GO) $(GOFLAG) get -v github.com/swaggo/gin-swagger
	$(GO) $(GOFLAG) get -v github.com/swaggo/files

swagger:
	@echo "Generates docs"
	swag init -g ./main.go --output docs/

generate:
	@echo "Generates mocks"
	mockery --all --keeptree

test:
	go test ./...

.PHONY: all install format lint swagger test
