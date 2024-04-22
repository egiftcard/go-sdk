package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/egiftcard/go-sdk/v2/egiftcard"
)

var (
	egiftcardSDK *egiftcard.EgiftcardSDK
)

func initSdk() {
	if sdk, err := egiftcard.NewEgiftcardSDK(
		chainRpcUrl,
		&egiftcard.SDKOptions{
			PrivateKey: privateKey,
		},
	); err != nil {
		panic(err)
	} else {
		egiftcardSDK = sdk
	}
}

func awaitTx(hash common.Hash) (*types.Transaction, error) {
	provider := egiftcardSDK.GetProvider()
	wait := time.Second * 1
	maxAttempts := uint8(20)
	attempts := uint8(0)

	var syncError error
	for {
		if attempts >= maxAttempts {
			fmt.Println("Retry attempts to get tx exhausted, tx might have failed")
			return nil, syncError
		}

		if tx, isPending, err := provider.TransactionByHash(context.Background(), hash); err != nil {
			syncError = err
			log.Printf("Failed to get tx %v, err = %v\n", hash.String(), err)
			attempts += 1
			time.Sleep(wait)
			continue
		} else {
			if isPending {
				log.Println("Transaction still pending...")
				time.Sleep(wait)
				continue
			}
			log.Printf("Transaction with hash %v mined successfully\n", tx.Hash())
			return tx, nil
		}
	}
}

func getNftCollection() (*egiftcard.NFTCollection, error) {
	if egiftcardSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Collection on chain %v, contract %v\n", chainRpcUrl, nftContractAddress)

	if contract, err := egiftcardSDK.GetNFTCollection(nftContractAddress); err != nil {
		log.Println("Failed to create an NFT Collection object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEdition() (*egiftcard.Edition, error) {
	if egiftcardSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition on chain %v, contract %v\n", chainRpcUrl, editionAddress)

	if contract, err := egiftcardSDK.GetEdition(editionAddress); err != nil {
		log.Println("Failed to create an Edition object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getToken() (*egiftcard.Token, error) {
	if egiftcardSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Token on chain %v, contract %v\n", chainRpcUrl, tokenAddress)

	if contract, err := egiftcardSDK.GetToken(tokenAddress); err != nil {
		log.Println("Failed to create an Token object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getNftDrop() (*egiftcard.NFTDrop, error) {
	if egiftcardSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Drop on chain %v, contract %v\n", chainRpcUrl, nftDropContractAddress)

	if contract, err := egiftcardSDK.GetNFTDrop(nftDropContractAddress); err != nil {
		log.Println("Failed to create an NFT Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEditionDrop() (*egiftcard.EditionDrop, error) {
	if egiftcardSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition Drop on chain %v, contract %v\n", chainRpcUrl, editionDropContractAddress)

	if contract, err := egiftcardSDK.GetEditionDrop(editionDropContractAddress); err != nil {
		log.Println("Failed to create an Edition Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMultiwrap() (*egiftcard.Multiwrap, error) {
	if egiftcardSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Multiwrap on chain %v, contract %v\n", chainRpcUrl, multiwrapContractAddress)

	if contract, err := egiftcardSDK.GetMultiwrap(multiwrapContractAddress); err != nil {
		log.Println("Failed to create a Multiwrap object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMarketplace() (*egiftcard.Marketplace, error) {
	if egiftcardSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Marketplace on chain %v, contract %v\n", chainRpcUrl, marketplaceAddress)

	if contract, err := egiftcardSDK.GetMarketplace(marketplaceAddress); err != nil {
		log.Println("Failed to create a Marketplace object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getCustom() (*egiftcard.SmartContract, error) {
	if egiftcardSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Custom on chain %v, contract %v\n", chainRpcUrl, customContractAddress)

	if contract, err := egiftcardSDK.GetContract(context.Background(), customContractAddress); err != nil {
		log.Println("Failed to create an Custom object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getStorage() egiftcard.IpfsStorage {
	if egiftcardSDK == nil {
		initSdk()
	}

	return egiftcardSDK.Storage
}
