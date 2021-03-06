# Absolute github repository name.
REPO := github.com/rigoiot/atlas-app-toolkit

# Build directory absolute path.
PROJECT_ROOT = $(CURDIR)

# Utility docker image to generate Go files from .proto definition.
# https://github.com/infobloxopen/atlas-gentool
GENTOOL_IMAGE := infoblox/atlas-gentool:latest

# Path for tools
BINDIR	:= bin

.PHONY: default
default: test

test: check-fmt vendor
	go test -cover ./...

.PHONY: vendor
vendor:
	dep ensure

check-fmt:
	test -z `go fmt ./...`

$(BINDIR)/gogofast:
	@echo "--> building $@"
	@go build -o $@ github.com/gogo/protobuf/protoc-gen-gogofast

.gen-query: $(BINDIR)/gogofast
	protoc -I. \
	-I$(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway) \
	--plugin=protoc-gen-gogofast=$(BINDIR)/gogofast --gogofast_out=paths=source_relative:. \
	query/collection_operators.proto

.gen-errdetails:
	docker run --rm -v $(PROJECT_ROOT):/go/src/$(REPO) $(GENTOOL_IMAGE) \
	--go_out=:. $(REPO)/rpc/errdetails/error_details.proto

.gen-errfields:
	docker run --rm -v $(PROJECT_ROOT):/go/src/$(REPO) $(GENTOOL_IMAGE) \
	--go_out=:. $(REPO)/rpc/errfields/error_fields.proto

.gen-servertestdata:
	docker run --rm -v $(PROJECT_ROOT):/go/src/$(REPO) $(GENTOOL_IMAGE) \
	--go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. $(REPO)/server/testdata/test.proto

.PHONY: gen
gen: .gen-query .gen-errdetails .gen-errfields
