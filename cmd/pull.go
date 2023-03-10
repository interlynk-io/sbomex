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

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Downloads specified SBOM from the repository and prints to the screen",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := logger.WithLogger(context.Background())
		processPull(ctx)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().Int32Var(&id, "id", 0, "pull SBOM based on Id (Ignores --filter)")
	pullCmd.MarkFlagRequired("id")
}

func processPull(ctx context.Context) {
	if isInValidCMD() {
		return
	}
	sbomlcDB, _ := db.NewSbomlc()
	url := sbomlcDB.Url(&model.CMDArgs{
		Id:   id,
	})
	if url != "" {
		http.Get(url)
	}
}
