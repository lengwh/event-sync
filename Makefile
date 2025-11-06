TM_ABI_ARTIFACT := ./abis/TreasureManager.sol/TreasureManager.json


event-sync:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/event-sync

clean:
	rm event-sync

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings:
	$(eval temp := $(shell mktemp))

	cat $(TM_ABI_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(TM_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/treasure_manager.go \
		--type TreasureManager \
		--bin $(temp)

		rm $(temp)

.PHONY: \
	event-sync \
	bindings \
	clean \
	test \
	lint