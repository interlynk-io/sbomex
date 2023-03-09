/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/interlynk-io/sbomex/pkg/http"

	"github.com/interlynk-io/sbomex/pkg/db"
	"github.com/interlynk-io/sbomex/pkg/logger"
	"github.com/interlynk-io/sbomex/pkg/model"
	"github.com/spf13/cobra"
)

var id int32
var filter string

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := logger.WithLogger(context.Background())
		processFetch(ctx)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().Int32Var(&id, "id", 0, "Fetch SBOM based on Id (Ignores --filter)")
	fetchCmd.Flags().StringVar(&filter, "filter", "", "Filter SBOM based on conditions provided")
}

func processFetch(ctx context.Context) {

	sbomlcDB, _ := db.NewSbomlc()

	http.Get(sbomlcDB.Url(&model.CMDArgs{
		Id:   id,
		Tool: filter,
	}))

}
