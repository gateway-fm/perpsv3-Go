# generate go files for all contracts
generate-contracts-all: generate-contracts-goerli generate-contracts-optimism

# generate go files for goerli net contracts
generate-contracts-goerli: generate-core-goerli generate-spot_market-goerli generate-perps_market-goerli

# generate fo giles for optimism net contracts
generate-contracts-optimism: generate-core-optimism generate-spot_market-optimism

# generate all mocks
mock-all: mock-service

# generate go file for SynthetixCore contract on goerli net
generate-core-goerli:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./Synthetix-Gitbook-v3/for-developers/abis/420-SynthetixCore.json ./contracts/coreGoerli
	abigen --abi=./contracts/420-SynthetixCore.json --pkg=coreGoerli --out=./contracts/coreGoerli/contract.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/420-SynthetixCore.json

# generate go file for SpotMarket contract on goerli net
generate-spot_market-goerli:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./Synthetix-Gitbook-v3/for-developers/abis/420-SpotMarket.json ./contracts/spotMarketGoerli
	abigen --abi=./contracts/420-SpotMarket.json --pkg=spotMarketGoerli --out=./contracts/spotMarketGoerli/contract.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/420-SpotMarket.json

# generate go file for PerpsMarket contract on goerli net
generate-perps_market-goerli:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./Synthetix-Gitbook-v3/for-developers/abis/420-PerpsMarket.json ./contracts/perpsMarketGoerli
	abigen --abi=./contracts/420-PerpsMarket.json --pkg=perpsMarketGoerli --out=./contracts/perpsMarketGoerli/contract.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/420-PerpsMarket.json

# generate go file for SynthetixCore contract on optimism net
generate-core-optimism:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./Synthetix-Gitbook-v3/for-developers/abis/10-SynthetixCore.json ./contracts/coreOptimism
	abigen --abi=./contracts/10-SynthetixCore.json --pkg=coreOptimism --out=./contracts/coreOptimism/contract.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/10-SynthetixCore.json

# generate go file for SpotMarket contract on optimism net
generate-spot_market-optimism:
	go run ./utils/getAbis/get-abis.go --get-mkdir ./Synthetix-Gitbook-v3/for-developers/abis/10-SpotMarket.json ./contracts/spotMarketOptimism
	abigen --abi=./contracts/10-SpotMarket.json --pkg=spotMarketOptimism --out=./contracts/spotMarketOptimism/contract.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/10-SpotMarket.json

# update Synthetix-Gitbook-v3 subtree
update-subtree:
	git subtree pull --prefix Synthetix-Gitbook-v3 git@github.com:Synthetixio/Synthetix-Gitbook-v3.git en --squash

# generate mock for service interface for testing
mock-service:
	mockgen -source=services/service.go -destination=mocks/service/mockService.go

tidy:
	go mod tidy

test: mock-all
	go test ./...

test-v: mock-all
	go test ./... -v

test-coverage: mock-all
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

lint:
	golint ./......

