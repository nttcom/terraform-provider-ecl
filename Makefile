TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -not -iwholename './.git/*' -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=ecl

default: build

build: fmtcheck
	go install

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

testacc-all: fmtcheck
	TF_ACC=1 go test ./ecl -v -count=1 -timeout 240m -parallel 4

testacc-short: fmtcheck
	TF_ACC=1 go test ./ecl -v -short -count=1 -timeout 120m -parallel 4

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

vendor:
	go mod vendor

.PHONY: build test testacc testacc-all testacc-short vet fmt fmtcheck errcheck test-compile vendor
