/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/interlynk-io/sbomex/pkg/db"
	"github.com/interlynk-io/sbomex/pkg/logger"
	"github.com/interlynk-io/sbomex/pkg/model"
	"github.com/interlynk-io/sbomex/pkg/view"

	"github.com/spf13/cobra"
)

var format string
var spec string
var tool string
var limit int32

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Finds SBOM in the repository that matches the filtering criteria",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := logger.WithLogger(context.Background())
		processSearch(ctx)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVar(&format, "format", "", "string SBOM format options json/xml/tv")
	searchCmd.Flags().StringVar(&spec, "spec", "", "string SBOM Specification options spdx/cdx")
	searchCmd.Flags().StringVar(&tool, "tool", "", "string SBOM creator tool name (e.g. syft, trivy, bom)")
	searchCmd.Flags().Int32Var(&limit, "limit", 25, "int max number of search results to print (default 25)")
}

func processSearch(ctx context.Context) {
	log := logger.FromContext(ctx)
	log.Debugf("Processing search")
	if isInValidCMD() {
		return
	}
	sbomlcDB, _ := db.NewSbomlc()
	view.SearchView(ctx, sbomlcDB.Search(&model.CMDArgs{
		Format: format,
		Spec:   spec,
		Tool:   tool,
		Limit:  limit,
	}))

}

func isInValidCMD() bool {
	if format != "" && format != "json" && format != "xml" && format != "tv" {
		fmt.Printf("invalid spec")
		return true
	}

	if spec != "" && spec != "spdx" && spec != "cdx" {
		fmt.Printf("invalid spec")
		return true
	}

	if id < 0 {
		fmt.Println("invalid id")
		return true
	}
	return false
}
