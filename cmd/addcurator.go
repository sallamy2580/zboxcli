package cmd

import (
	"log"

	"github.com/0chain/gosdk/zboxcore/sdk"
	"github.com/spf13/cobra"
)

var addCuratorCmd = &cobra.Command{
	Use:   "addcurator",
	Short: "Adds a curator to an allocation",
	Long:  "Adds a curator to an allocation",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var flags = cmd.Flags()

		if !flags.Changed("allocation") {
			log.Fatal("Error: allocation flag is missing")
		}
		allocationID, err := flags.GetString("allocation")
		if err != nil {
			log.Fatal("invalid 'allocation' flag: ", err)
		}

		if !flags.Changed("curator") {
			log.Fatal("Error: curator flag is missing")
		}
		curatorID, err := flags.GetString("curator")
		if err != nil {
			log.Fatal("invalid 'curator' flag: ", err)
		}

		_, _, err = sdk.AddCurator(curatorID, allocationID)

		if err != nil {
			log.Fatal("Error adding curator:", err)
		}
		log.Println(clientWallet.ClientID + " added " + curatorID + " as a curator to allocation " + allocationID)
	},
}

func init() {
	rootCmd.AddCommand(addCuratorCmd)
	addCuratorCmd.PersistentFlags().
		String("curator", "",
			"new curator to add to allocation")
	addCuratorCmd.PersistentFlags().
		String("allocation", "",
			"allocation that the curator is to be added")

	addCuratorCmd.MarkFlagRequired("curator")
	addCuratorCmd.MarkFlagRequired("allocation")
}
