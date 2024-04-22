
## Contract Deployments

The contract deployer lets you deploy new contracts to the blockchain using just the egiftcard SDK. You can access the contract deployer interface as follows:

```
import (
	"github.com/egiftcard/go-sdk/v2/egiftcard"
)

privateKey := "..."
secretKey := "..."

sdk, err := egiftcard.NewEgiftcardSDK("mumbai", &egiftcard.SDKOptions{
	PrivateKey: privateKey,
	SecretKey: secretKey
})

// Now you can deploy a contract
address, err := sdk.Deployer.DeployNFTCollection(
	&egiftcard.DeployNFTCollectionMetadata{
		Name: "Go NFT",
	}
})
```

```go
type ContractDeployer struct {
    *ProviderHandler
}
```

### func \(\*ContractDeployer\) [DeployEdition](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/contract_deployer.go#L117>)

```go
func (deployer *ContractDeployer) DeployEdition(ctx context.Context, metadata *DeployEditionMetadata) (string, error)
```

Deploy a new Edition contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployEdition(
     context.Background(),
		&egiftcard.DeployEditionMetadata{
			Name: "Go Edition",
		}
	})
```

### func \(\*ContractDeployer\) [DeployEditionDrop](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/contract_deployer.go#L174>)

```go
func (deployer *ContractDeployer) DeployEditionDrop(ctx context.Context, metadata *DeployEditionDropMetadata) (string, error)
```

Deploy a new Edition Drop contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployEditionDrop(
     context.Background(),
		&egiftcard.DeployEditionDropMetadata{
			Name: "Go Edition Drop",
		}
	})
```

### func \(\*ContractDeployer\) [DeployMarketplace](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/contract_deployer.go#L212>)

```go
func (deployer *ContractDeployer) DeployMarketplace(ctx context.Context, metadata *DeployMarketplaceMetadata) (string, error)
```

Deploy a new Marketplace contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployMarketplace(
     context.Background()
		&egiftcard.DeployMarketplaceMetadata{
			Name: "Go Marketplace",
		}
	})
```

### func \(\*ContractDeployer\) [DeployMultiwrap](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/contract_deployer.go#L193>)

```go
func (deployer *ContractDeployer) DeployMultiwrap(ctx context.Context, metadata *DeployMultiwrapMetadata) (string, error)
```

Deploy a new Multiwrap contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployMultiwrap(
     context.Background()
		&egiftcard.DeployMultiwrapMetadata{
			Name: "Go Multiwrap",
		}
	})
```

### func \(\*ContractDeployer\) [DeployNFTCollection](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/contract_deployer.go#L98>)

```go
func (deployer *ContractDeployer) DeployNFTCollection(ctx context.Context, metadata *DeployNFTCollectionMetadata) (string, error)
```

Deploy a new NFT Collection contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployNFTCollection(
     context.Background(),
		&egiftcard.DeployNFTCollectionMetadata{
			Name: "Go NFT",
		}
	})
```

### func \(\*ContractDeployer\) [DeployNFTDrop](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/contract_deployer.go#L155>)

```go
func (deployer *ContractDeployer) DeployNFTDrop(ctx context.Context, metadata *DeployNFTDropMetadata) (string, error)
```

Deploy a new NFT Drop contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployNFTDrop(
     context.Background(),
		&egiftcard.DeployNFTDropMetadata{
			Name: "Go NFT Drop",
		}
	})
```

### func \(\*ContractDeployer\) [DeployToken](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/contract_deployer.go#L136>)

```go
func (deployer *ContractDeployer) DeployToken(ctx context.Context, metadata *DeployTokenMetadata) (string, error)
```

Deploy a new Token contract.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployToken(
     context.Background(),
		&egiftcard.DeployTokenMetadata{
			Name: "Go Token",
		}
	})
```
