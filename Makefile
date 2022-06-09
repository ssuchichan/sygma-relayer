
.PHONY: help run build install license
all: help

## license: Adds license header to missing files.
license:
	@echo "  >  \033[32mAdding license headers...\033[0m "
	GO111MODULE=off go get -u github.com/google/addlicense
	addlicense -c "ChainSafe Systems" -f ./scripts/header.txt -y 2021 .

## license-check: Checks for missing license headers
license-check:
	@echo "  >  \033[Checking for license headers...\033[0m "
	GO111MODULE=off go get -u github.com/google/addlicense
	addlicense -check -c "ChainSafe Systems" -f ./scripts/header.txt -y 2021 .


coverage:
	go tool cover -func cover.out | grep total | awk '{print $3}'

test:
	./scripts/tests.sh

## Install dependency subkey
install-subkey:
	curl https://getsubstrate.io -sSf | bash -s -- --fast
	cargo install --force --git https://github.com/paritytech/substrate subkey

genmocks:
	mockgen -destination=./chains/evm/calls/evmgaspricer/mock/gas-pricer.go -source=./chains/evm/calls/evmgaspricer/gas-pricer.go
	mockgen -destination=./relayer/mock/relayer.go -source=./relayer/relayer.go
	mockgen -source=chains/evm/calls/calls.go -destination=chains/evm/calls/mock/calls.go
	mockgen -source=chains/evm/calls/transactor/transact.go -destination=chains/evm/calls/transactor/mock/transact.go
	mockgen -destination=./chains/evm/calls/transactor/itx/mock/itx.go -source=./chains/evm/calls/transactor/itx/itx.go
	mockgen -destination=./chains/evm/calls/transactor/itx//mock/minimalForwarder.go -source=./chains/evm/calls/transactor/itx/minimalForwarder.go
	mockgen -destination=./tss/common/mock/tss.go github.com/binance-chain/tss-lib/tss Message
	mockgen -destination=./tss/common/mock/communication.go -source=./tss/common/base.go -package mock_tss
	mockgen -destination=./tss/keygen/mock/storer.go -source=./tss/keygen/keygen.go
	mockgen -destination=./tss/keygen/mock/storer.go -source=./tss/keygen/keygen.go
	mockgen -source=./tss/coordinator.go -destination=./tss/mock/coordinator.go
	mockgen -source=./comm/communication.go -destination=./comm/mock/communication.go

e2e-test:
	./scripts/e2e_tests.sh

example:
	docker-compose --file=./example/docker-compose.yml up