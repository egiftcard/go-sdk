
## Edition Drop

You can access the Edition Drop interface from the SDK as follows:

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

contract, err := sdk.GetEditionDrop("{{contract_address}}")
```

```go
type EditionDrop struct {
    *ERC1155Standard

    Helper          *contractHelper
    ClaimConditions *EditionDropClaimConditions
    Encoder         *ContractEncoder
    Events          *ContractEvents
}
```

### func \(\*EditionDrop\) [Claim](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/edition_drop.go#L121>)

```go
func (drop *EditionDrop) Claim(ctx context.Context, tokenId int, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet.

tokenId: the token ID of the NFT to claim

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*EditionDrop\) [ClaimTo](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/edition_drop.go#L142>)

```go
func (drop *EditionDrop) ClaimTo(ctx context.Context, destinationAddress string, tokenId int, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet.

tokenId: the token ID of the NFT to claim

destinationAddress: the address of the wallet to claim the NFTs to

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

#### Example

```
address = "{{wallet_address}}"
tokenId = 0
quantity = 1

tx, err := contract.ClaimTo(context.Background(), address, tokenId, quantity)
```

### func \(\*EditionDrop\) [CreateBatch](<https://github.com/egiftcard/go-sdk/blob/main/egiftcard/edition_drop.go#L110>)

```go
func (drop *EditionDrop) CreateBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Create a batch of NFTs on this contract.

metadatas: a list of the metadatas of the NFTs to create

returns: the transaction receipt of the batch creation

#### Example

```
image0, err := os.Open("path/to/image/0.jpg")
defer image0.Close()

image1, err := os.Open("path/to/image/1.jpg")
defer image1.Close()

metadatasWithSupply := []*egiftcard.EditionMetadataInput{
	&egiftcard.EditionMetadataInput{
		Metadata: &egiftcard.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
			Image: image0,
		},
		Supply: 100,
	},
	&egiftcard.EditionMetadataInput{
		Metadata: &egiftcard.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
			Image: image1,
		},
		Supply: 100,
	},
}

tx, err := contract.MintBatchTo(context.Background(), "{{wallet_address}}", metadatasWithSupply)
```
