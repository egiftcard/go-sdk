.PHONY: abi docs publish local-test

export EGIFTCARD_SECRET_KEY

SHELL := /bin/bash

abi:
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC20.json --out abi/token_erc20.go --type TokenERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC721.json --out abi/token_erc721.go --type TokenERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TokenERC1155.json --out abi/token_erc1155.go --type TokenERC1155
	# If you want to generate drop contracts, you'll have to delete a struct
	# abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC721_V3.json --out abi/drop_erc721.go --type DropERC721
	# abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/DropERC1155_V2.json --out abi/drop_erc1155.go --type DropERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/Multiwrap.json --out abi/multiwrap.go --type Multiwrap
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/Marketplace.json --out abi/marketplace.go --type Marketplace

	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/TWFactory.json --out abi/twfactory.go --type TWFactory
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC20.json --out abi/ierc20.go --type IERC20
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC721.json --out abi/ierc721.go --type IERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC1155.json --out abi/ierc1155.go --type IERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IERC165.json --out abi/ierc165.go --type IERC165
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/SignatureMintERC721.json --out abi/signature_mint_erc721.go --type SignatureMintERC721
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/SignatureMintERC1155.json --out abi/signature_mint_erc1155.go --type SignatureMintERC1155
	abigen --alias contractURI=internalContractURI --pkg abi --abi internal/json/IRoyalty.json --out abi/iroyalty.go --type IRoyalty

docs:
	rm -rf docs
	mkdir docs
	gomarkdoc --output docs/doc.md --repository.default-branch main ./egiftcard
	node ./scripts/generate-docs.mjs
	rm ./docs/doc.md ./docs/start.md ./docs/delete.md
	node ./scripts/generate-snippets.mjs

cmd: FORCE
	cd cmd/egiftcard && go build -o ../../bin/egiftcard && cd -

test-nft-read:
	./bin/egiftcard nft getAll -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard nft getOwned -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-write:
	./bin/egiftcard nft mint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard nft mintLink -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-nft-sigmint:
	./bin/egiftcard nft sigmint -a ${GO_NFT_COLLECTION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-read:
	./bin/egiftcard edition getAll -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard edition getOwned -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-write:
	./bin/egiftcard edition mint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-sigmint:
	./bin/egiftcard edition sigmint -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard edition sigmint-tokenid -a ${GO_EDITION} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-read:
	./bin/egiftcard token get -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-token-write:
	./bin/egiftcard token mint -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard token mintBatch -a ${GO_TOKEN} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-multiwrap-read:
	./bin/egiftcard multiwrap getAll -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	# ./bin/egiftcard multiwrap getContents -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-multiwrap-write:
	./bin/egiftcard multiwrap wrap -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC} -n ${GO_NFT_COLLECTION} -e ${GO_EDITION} -t ${GO_TOKEN}
	./bin/egiftcard multiwrap unwrap -a ${GO_MULTIWRAP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-read:
	./bin/egiftcard nftdrop getAll -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard nftdrop getActive -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-drop-write:
	./bin/egiftcard nftdrop createBatch -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard nftdrop claim -a ${GO_NFT_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-drop-read:
	./bin/egiftcard editiondrop getAll -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-edition-drop-write:
	./bin/egiftcard editiondrop createBatch -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard editiondrop claim -a ${GO_EDITION_DROP} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-storage:
	./bin/egiftcard storage upload
	./bin/egiftcard storage uploadBatch
	./bin/egiftcard storage uploadImage
	./bin/egiftcard storage uploadImageLink

test-custom:
	./bin/egiftcard custom set -a ${GO_CUSTOM} -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-deploy:
	./bin/egiftcard deploy nft -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard deploy edition -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard deploy token -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard deploy nftdrop -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard deploy editiondrop -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard deploy multiwrap -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}
	./bin/egiftcard deploy marketplace -k ${GO_PRIVATE_KEY} -u ${GO_ALCHEMY_RPC}

test-encoder:
	./bin/egiftcard marketplace encode -a ${GO_MARKETPLACE} -k ${GO_PRIVATE_KEY} -u ${GO_AVAX_RPC}

test-cmd:
	make cmd
	make test-nft-read
	make test-nft-write
	make test-edition-read
	make test-edition-write
	make test-token-read
	make test-token-write
	make test-drop-read
	make test-drop-write
	make test-edition-drop-read
	make test-edition-drop-write
	make test-multiwrap-read
	make test-multiwrap-write
	make test-storage

stop-docker:
	docker stop hardhat-node
	docker rm hardhat-node

test: FORCE
	docker build . -t hardhat-mainnet-fork
	docker start hardhat-node || docker run --name hardhat-node -d -p 8545:8545 -e "EGIFTCARD_SECRET_KEY=${EGIFTCARD_SECRET_KEY}" -e "SDK_ALCHEMY_KEY=${SDK_ALCHEMY_KEY}" hardhat-mainnet-fork
	sudo bash ./scripts/test/await-hardhat.sh
	go clean -testcache
	go test -v ./egiftcard
	docker stop hardhat-node
	docker rm hardhat-node

local-test:
  # Needs to be run along with npx hardhat node from this repo, and needs to be a mainnet fork hardhat
	go clean -testcache
	go test -v ./egiftcard

publish:
	# Make sure to pass the TAG variable to this command ex: `make publish TAG=v2.0.0`
	# Publish following https://go.dev/doc/modules/publishing
	go mod tidy
	git tag $(TAG)
	git push origin $(TAG)
	GOPROXY=proxy.golang.org go list -m github.com/egiftcard/go-sdk/v2@$(TAG)

FORCE: ;