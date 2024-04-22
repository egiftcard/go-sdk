
## EgiftcardSDK

```go
type EgiftcardSDK struct {
    *ProviderHandler
    Storage  IpfsStorage
    Deployer ContractDeployer
    Auth     WalletAuthenticator
}
```

### func [NewEgiftcardSDK](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L27>)

```go
func NewEgiftcardSDK(rpcUrlOrChainName string, options *SDKOptions) (*EgiftcardSDK, error)
```

#### NewEgiftcardSDK

\# Create a new instance of the Egiftcard SDK

rpcUrlOrName: the name of the chain to connection to \(e.g. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche"\) or the RPC URL to connect to

options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL

### func [NewEgiftcardSDKFromProvider](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L46>)

```go
func NewEgiftcardSDKFromProvider(provider *ethclient.Client, options *SDKOptions) (*EgiftcardSDK, error)
```

### func \(\*EgiftcardSDK\) [GetContract](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L213>)

```go
func (sdk *EgiftcardSDK) GetContract(ctx context.Context, address string) (*SmartContract, error)
```

#### GetContract

\# Get an instance of a custom contract deployed with egiftcard deploy

address: the address of the contract

### func \(\*EgiftcardSDK\) [GetContractFromAbi](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L229>)

```go
func (sdk *EgiftcardSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error)
```

#### GetContractFromABI

\# Get an instance of ant custom contract from its ABI

address: the address of the contract

abi: the ABI of the contract

### func \(\*EgiftcardSDK\) [GetEdition](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L127>)

```go
func (sdk *EgiftcardSDK) GetEdition(address string) (*Edition, error)
```

#### GetEdition

\# Get an Edition contract SDK instance

address: the address of the Edition contract

### func \(\*EgiftcardSDK\) [GetEditionDrop](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L171>)

```go
func (sdk *EgiftcardSDK) GetEditionDrop(address string) (*EditionDrop, error)
```

#### GetEditionDrop

\# Get an Edition Drop contract SDK instance

address: the address of the Edition Drop contract

### func \(\*EgiftcardSDK\) [GetMarketplace](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L199>)

```go
func (sdk *EgiftcardSDK) GetMarketplace(address string) (*Marketplace, error)
```

#### GetMarketplace

\# Get a Marketplace contract SDK instance

address: the address of the Marketplace contract

### func \(\*EgiftcardSDK\) [GetMultiwrap](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L185>)

```go
func (sdk *EgiftcardSDK) GetMultiwrap(address string) (*Multiwrap, error)
```

#### GetMultiwrap

\# Get a Multiwrap contract SDK instance

address: the address of the Multiwrap contract

### func \(\*EgiftcardSDK\) [GetNFTCollection](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L113>)

```go
func (sdk *EgiftcardSDK) GetNFTCollection(address string) (*NFTCollection, error)
```

#### GetNFTCollection

\# Get an NFT Collection contract SDK instance

address: the address of the NFT Collection contract

### func \(\*EgiftcardSDK\) [GetNFTDrop](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L157>)

```go
func (sdk *EgiftcardSDK) GetNFTDrop(address string) (*NFTDrop, error)
```

#### GetNFTDrop

\# Get an NFT Drop contract SDK instance

address: the address of the NFT Drop contract

### func \(\*EgiftcardSDK\) [GetToken](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/sdk.go#L143>)

```go
func (sdk *EgiftcardSDK) GetToken(address string) (*Token, error)
```

#### GetToken

\# Returns a Token contract SDK instance

address: address of the token contract

Returns a Token contract SDK instance
