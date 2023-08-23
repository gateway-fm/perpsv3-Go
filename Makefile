generate-contracts-all: generate-contracts-goerli generate-contracts-optimism
generate-contracts-goerli: generate-core-goerli generate-spot_market-goerli generate-preps_market-goerli
generate-contracts-optimism: generate-core-optimism generate-spot_market-optimism

generate-core-goerli:
	go run ./utils/getAbis/get-abis.go --get ./Synthetix-Gitbook-v3/for-developers/abis/420-SynthetixCore.json
	abigen --abi=./contracts/420-SynthetixCore.json --pkg=contracts --out=./contracts/core-goerli.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/420-SynthetixCore.json

generate-spot_market-goerli:
	go run ./utils/getAbis/get-abis.go --get ./Synthetix-Gitbook-v3/for-developers/abis/420-SpotMarket.json
	abigen --abi=./contracts/420-SpotMarket.json --pkg=contracts --out=./contracts/spot-market-goerli.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/420-SpotMarket.json

generate-preps_market-goerli:
	go run ./utils/getAbis/get-abis.go --get ./Synthetix-Gitbook-v3/for-developers/abis/420-PerpsMarket.json
	abigen --abi=./contracts/420-PerpsMarket.json --pkg=contracts --out=./contracts/preps-market-goerli.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/420-PerpsMarket.json

generate-core-optimism:
	go run ./utils/getAbis/get-abis.go --get ./Synthetix-Gitbook-v3/for-developers/abis/10-SynthetixCore.json
	abigen --abi=./contracts/10-SynthetixCore.json --pkg=contracts --out=./contracts/core-optimism.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/10-SynthetixCore.json

generate-spot_market-optimism:
	go run ./utils/getAbis/get-abis.go --get ./Synthetix-Gitbook-v3/for-developers/abis/10-SpotMarket.json
	abigen --abi=./contracts/10-SpotMarket.json --pkg=contracts --out=./contracts/spot-market-optimism.go
	go run ./utils/getAbis/get-abis.go --rm ./contracts/10-SpotMarket.json