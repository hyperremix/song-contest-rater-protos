all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## go-proto-generate: generate protobuf files
.PHONY: go-proto-generate
go-proto-generate:
	protoc -I=. --go_out=. --go-grpc_out=. ./*.proto
