# generate go files for goerli net contracts
generate-contracts-andromeda: generate-core-andromeda generate-perps_market-andromeda generate-susdt-andromeda generate-forwarder-andromeda generate-erc7412-andromeda

# generate all mocks
mock-all: mock-service mock-events

# generate go file for SynthetixCore contract on andromeda net
generate-core-andromeda:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./Synthetix-Gitbook-v3/for-developers/abis/84531-andromeda-SynthetixCore.json ./contracts/core
	abigen --abi=./contracts/84531-andromeda-SynthetixCore.json --pkg=core --out=./contracts/core/contract.go

# generate go file for PerpsMarket contract on andromeda net
generate-perps_market-andromeda:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./Synthetix-Gitbook-v3/for-developers/abis/84531-andromeda-PerpsMarket.json ./contracts/perpsMarket
	abigen --abi=./contracts/84531-andromeda-PerpsMarket.json --pkg=perpsMarket --out=./contracts/perpsMarket/contract.go

# generate go file for PerpsMarket contract on andromeda net from cannon abi
generate-perps_market-andromeda-c:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./cannon-synthetix/base/perpsFactory/PerpsMarketProxy.json ./contracts/perpsMarket
	abigen --abi=./contracts/PerpsMarketProxy.json --pkg=perpsMarket --out=./contracts/perpsMarket/contract.go

# generate go file for snxUSDT contract on andromeda net
generate-susdt-andromeda:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./Synthetix-Gitbook-v3/for-developers/abis/84531-andromeda-snxUSDToken.json ./contracts/sUSDT
	abigen --abi=./contracts/84531-andromeda-snxUSDToken.json --pkg=sUSDT --out=./contracts/sUSDT/contract.go

# generate go file for TrustedMulticallForwarder contract on andromeda net
generate-forwarder-andromeda:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./additional-abis/84531-TrustedMulticallForwarder.json ./contracts/forwarder
	abigen --abi=./contracts/84531-TrustedMulticallForwarder.json --pkg=forwarder --out=./contracts/forwarder/contract.go

generate-forwarder-mainnet-c:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./cannon-synthetix/base/system/trusted_multicall_forwarder/TrustedMulticallForwarder.json ./contracts/forwarder
	abigen --abi=./contracts/TrustedMulticallForwarder.json --pkg=forwarder --out=./contracts/forwarder/contract.go


# generate go file for ERC7412 contract on andromeda net
generate-erc7412-andromeda:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./additional-abis/84531-ERC7412.json ./contracts/ERC7412
	abigen --abi=./contracts/84531-ERC7412.json --pkg=erc7412 --out=./contracts/ERC7412/contract.go

generate-erc7412-mainnet-c:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./cannon-synthetix/base/pyth_erc7412_wrapper/PythERC7412Wrapper.json ./contracts/ERC7412
	abigen --abi=./contracts/PythERC7412Wrapper.json --pkg=ERC7412 --out=./contracts/ERC7412/contract.go

# update Synthetix-Gitbook-v3 subtree
update-subtree:
	git subtree pull --prefix Synthetix-Gitbook-v3 git@github.com:Synthetixio/Synthetix-Gitbook-v3.git en --squash

# fetch ABIs from cannon
fetch-cannon-andromeda:
	cannon inspect synthetix-omnibus:latest@andromeda --chain-id 84531 -w ./cannon-synthetix/andromeda --sources

# fetch ABIs from cannon
fetch-cannon-base:
	cannon inspect synthetix-omnibus:latest@andromeda --chain-id 8453 -w ./cannon-synthetix/base --sources

# generate mock for service interface for testing
mock-service:
	mockgen -source=services/service.go -destination=mocks/service/mockService.go

# generate mock for events interface for testing
mock-events:
	mockgen -source=events/events.go -destination=mocks/events/mockEvents.go

tidy:
	go mod tidy

test: mock-all
	go test --short

test-v: mock-all
	go test --short -v

test-coverage: mock-all
	go test -coverprofile=coverage.out ./... --short
	go tool cover -html=coverage.out

lint:
	golint ./...

