// Copyright 2023 Interlynk.io
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"

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
	Short: "Pulls specified SBOM from the repository and prints to the screen",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := logger.WithLogger(context.Background())
		return processPull(ctx)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().Int32Var(&id, "id", 0, "Pull SBOM based on Id")
	pullCmd.MarkFlagRequired("id")
}

func processPull(ctx context.Context) error {
	if isInValidCMD() {
		return fmt.Errorf("invalid command")
	}
	sbomlcDB, _ := db.NewSbomlc()
	url := sbomlcDB.Url(&model.CMDArgs{
		Id: id,
	})
	if url != "" {
		http.Get(url)
	}

	return nil
}
