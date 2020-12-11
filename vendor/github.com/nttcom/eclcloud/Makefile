fmt:
	gofmt -s -w .

fmtcheck:
	(! gofmt -s -d . | grep '^')

vet:
	go vet ./...

test:
	go test ./... -count=1

.PHONY: fmt fmtcheck vet test
