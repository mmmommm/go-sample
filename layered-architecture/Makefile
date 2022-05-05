GOPATH = $(shell go env GOPATH)
WIRE = "$(shell ls ${GOPATH}/bin | grep wire"
GOMOCK = "$(shell ls ${GOPATH}/bin | grep gomock)"

.PHONY: test
test:
	go test $(FLAGS) ./...

.PHONY generate
generate:
	go generate -x ./...

.PHONY: wire
wire:
ifneq (wire, $(shell echo $(WIRE)))
	go install github.com/google/wire/cmd/wire@latest
endif
	$(GOPATH)/bin/wire