package egiftcard

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EgiftcardSDK struct {
	*ProviderHandler
	Deployer *ContractDeployer
	Storage  IpfsStorage
	Auth     WalletAuthenticator
}

// NewEgiftcardSDK
//
// # Create a new instance of the Egiftcard SDK
//
// rpcUrlOrName: the name of the chain to connection to (e.g. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche") or the RPC URL to connect to
//
// options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL
func NewEgiftcardSDK(rpcUrlOrChainName string, options *SDKOptions) (*EgiftcardSDK, error) {
	clientId := "" 
	if (options != nil && options.SecretKey != "") {
		clientId = deriveClientId(options.SecretKey)
	}

	rpc, err := getDefaultRpcUrl(rpcUrlOrChainName, clientId)
	if err != nil {
		return nil, err
	}

	provider, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	return NewEgiftcardSDKFromProvider(provider, options)
}

func NewEgiftcardSDKFromProvider(provider *ethclient.Client, options *SDKOptions) (*EgiftcardSDK, error) {
	// Define defaults for all the options
	secretKey := ""
	privateKey := ""
	gatewayUrl := defaultIpfsGatewayUrl
	httpClient := http.DefaultClient

	// Override defaults with the options that are defined
	if options != nil {
		if options.SecretKey != "" {
			secretKey = options.SecretKey
		}

		if options.PrivateKey != "" {
			privateKey = options.PrivateKey
		}

		if options.GatewayUrl != "" {
			gatewayUrl = options.GatewayUrl
		} else if secretKey != "" {
			clientId := deriveClientId(secretKey)
			gatewayUrl = fmt.Sprintf("https://%s.ipfscdn.io/ipfs/", clientId)
		} else {
			gatewayUrl = defaultIpfsGatewayUrl
		}

		if options.HttpClient != nil {
			httpClient = options.HttpClient
		}
	}

	storage := newIpfsStorage(secretKey, gatewayUrl, httpClient)

	handler, err := NewProviderHandler(provider, privateKey)
	if err != nil {
		return nil, err
	}

	deployer, err := newContractDeployer(provider, privateKey, storage)
	if err != nil {
		return nil, err
	}

	if deployer == nil {
		fmt.Println("Warning: Contract deployments are not supported on this network. SDK instantiated without a Deployer.")
	}

	auth, err := newWalletAuthenticator(provider, privateKey)
	if err != nil {
		return nil, err
	}

	sdk := &EgiftcardSDK{
		ProviderHandler: handler,
		Deployer:        deployer,
		Storage:         *storage,
		Auth:            *auth,
	}

	return sdk, nil
}

// GetNFTCollection
//
// # Get an NFT Collection contract SDK instance
//
// address: the address of the NFT Collection contract
func (sdk *EgiftcardSDK) GetNFTCollection(address string) (*NFTCollection, error) {
	return newNFTCollection(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetEdition
//
// # Get an Edition contract SDK instance
//
// address: the address of the Edition contract
func (sdk *EgiftcardSDK) GetEdition(address string) (*Edition, error) {
	return newEdition(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetToken
//
// # Returns a Token contract SDK instance
//
// address: address of the token contract
//
// Returns a Token contract SDK instance
func (sdk *EgiftcardSDK) GetToken(address string) (*Token, error) {
	return newToken(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetNFTDrop
//
// # Get an NFT Drop contract SDK instance
//
// address: the address of the NFT Drop contract
func (sdk *EgiftcardSDK) GetNFTDrop(address string) (*NFTDrop, error) {
	return newNFTDrop(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetEditionDrop
//
// # Get an Edition Drop contract SDK instance
//
// address: the address of the Edition Drop contract
func (sdk *EgiftcardSDK) GetEditionDrop(address string) (*EditionDrop, error) {
	return newEditionDrop(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetMultiwrap
//
// # Get a Multiwrap contract SDK instance
//
// address: the address of the Multiwrap contract
func (sdk *EgiftcardSDK) GetMultiwrap(address string) (*Multiwrap, error) {
	return newMultiwrap(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetMarketplace
//
// # Get a Marketplace contract SDK instance
//
// address: the address of the Marketplace contract
func (sdk *EgiftcardSDK) GetMarketplace(address string) (*Marketplace, error) {
	return newMarketplace(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetContract
//
// # Get an instance of a custom contract deployed with egiftcard deploy
//
// address: the address of the contract
func (sdk *EgiftcardSDK) GetContract(ctx context.Context, address string) (*SmartContract, error) {
	abi, err := fetchContractMetadataFromAddress(ctx, address, sdk.GetProvider(), &sdk.Storage)
	if err != nil {
		return nil, err
	}

	return sdk.GetContractFromAbi(address, abi)
}

// GetContractFromABI
//
// # Get an instance of ant custom contract from its ABI
//
// address: the address of the contract
//
// abi: the ABI of the contract
func (sdk *EgiftcardSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error) {
	return newSmartContract(
		sdk.GetProvider(),
		common.HexToAddress(address),
		abi,
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

func defaultRpc(network string, clientId string) (string, error) {
	defaultApiKey := "718c5c811c7f3224efb283e04faab56a8a5cbde78d92a6d4cb905b41985d3856"
	if clientId != "" {
		return fmt.Sprintf("https://%s.rpc.egiftcard.cc/%s", network, clientId), nil
	}

	return fmt.Sprintf("https://%s.rpc.egiftcard.cc/%s", network, defaultApiKey), nil
}

func getDefaultRpcUrl(rpcUrlorName string, clientId string) (string, error) {
	switch rpcUrlorName {
	case "mumbai":
		return defaultRpc("mumbai", clientId)
	case "goerli":
		return defaultRpc("goerli", clientId)
	case "polygon":
		return defaultRpc("polygon", clientId)
	case "mainnet", "ethereum":
		return defaultRpc("ethereum", clientId)
	case "fantom":
		return defaultRpc("fantom", clientId)
	case "avalanche":
		return defaultRpc("avalanche", clientId)
	case "optimism":
		return defaultRpc("optimism", clientId)
	case "optimism-goerli":
		return defaultRpc("optimism-goerli", clientId)
	case "arbitrum":
		return defaultRpc("arbitrum", clientId)
	case "arbitrum-goerli":
		return defaultRpc("arbitrum-goerli", clientId)
	default:
		if strings.HasPrefix(rpcUrlorName, "http") || strings.HasPrefix(rpcUrlorName, "wss") {
			return rpcUrlorName, nil
		} else {
			return defaultRpc(rpcUrlorName, clientId)
		}
	}
}
