lint:
	@echo Running linter
	golangci-lint cache clean
	golangci-lint run --config=./.golangci.yml