default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

# Generate docs
generate:
	cd tools; go generate ./..

# See https://golangci-lint.run/
lint:
	golangci-lint run

fmt:
	go fmt ./...
