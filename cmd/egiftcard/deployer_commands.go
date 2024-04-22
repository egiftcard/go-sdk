package main

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/egiftcard/go-sdk/v2/egiftcard"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [command]",
	Short: "Deploy a contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var deployNftCmd = &cobra.Command{
	Use:   "nft",
	Short: "Deploy an nft collection",
	Run: func(cmd *cobra.Command, args []string) {
		if egiftcardSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := egiftcardSDK.Deployer.DeployNFTCollection(context.Background(), &egiftcard.DeployNFTCollectionMetadata{
			Name: "Goku NFT",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployEditionCmd = &cobra.Command{
	Use:   "edition",
	Short: "Deploy an edition",
	Run: func(cmd *cobra.Command, args []string) {
		if egiftcardSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := egiftcardSDK.Deployer.DeployEdition(context.Background(), &egiftcard.DeployEditionMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Deploy an token",
	Run: func(cmd *cobra.Command, args []string) {
		if egiftcardSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := egiftcardSDK.Deployer.DeployToken(context.Background(), &egiftcard.DeployTokenMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployNFTDropCmd = &cobra.Command{
	Use:   "nftdrop",
	Short: "Deploy an nft drop",
	Run: func(cmd *cobra.Command, args []string) {
		if egiftcardSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := egiftcardSDK.Deployer.DeployNFTDrop(context.Background(), &egiftcard.DeployNFTDropMetadata{
			Name: "Go Script Drop",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployEditionDropCmd = &cobra.Command{
	Use:   "editiondrop",
	Short: "Deploy an edition drop",
	Run: func(cmd *cobra.Command, args []string) {
		if egiftcardSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := egiftcardSDK.Deployer.DeployEditionDrop(context.Background(), &egiftcard.DeployEditionDropMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployMultiwrapCmd = &cobra.Command{
	Use:   "multiwrap",
	Short: "Deploy a multiwrap",
	Run: func(cmd *cobra.Command, args []string) {
		if egiftcardSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := egiftcardSDK.Deployer.DeployMultiwrap(context.Background(), &egiftcard.DeployMultiwrapMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployMarketplaceCmd = &cobra.Command{
	Use:   "marketplace",
	Short: "Deploy a marketplace",
	Run: func(cmd *cobra.Command, args []string) {
		if egiftcardSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := egiftcardSDK.Deployer.DeployMarketplace(context.Background(), &egiftcard.DeployMarketplaceMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

func init() {
	deployCmd.AddCommand(deployNftCmd)
	deployCmd.AddCommand(deployEditionCmd)
	deployCmd.AddCommand(deployTokenCmd)
	deployCmd.AddCommand(deployNFTDropCmd)
	deployCmd.AddCommand(deployEditionDropCmd)
	deployCmd.AddCommand(deployMultiwrapCmd)
	deployCmd.AddCommand(deployMarketplaceCmd)
}
