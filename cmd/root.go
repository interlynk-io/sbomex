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
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-github/v52/github"
	"github.com/interlynk-io/sbomex/pkg/model"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	version "sigs.k8s.io/release-utils/version"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sbomex",
	Short: "Find & pull public SBOMs from Interlynk's SBOM repository",
	Long: `SBOM Explorer (sbomex) is a command line utility
to help query and fetch Interlynk's public SBOM repository.
The tool is intended to help familiarize with the specifications
and formats of common SBOM standards and the quality of produced SBOMs (See sbomqs).
The underlying repository is updated periodically with SBOMs
from a variety of sources built with many tools`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	checkIfLatestRelease()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	downloadDB(model.SbomlcDataSource, model.DbLocation)
}

func downloadDB(path string, url string) {
	var out *os.File
	_, err := os.ReadFile(path)
	if err != nil {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			fmt.Printf("failed to create directory %s", err.Error())
			return
		}
		out, err = os.Create(path)
		if err != nil {
			fmt.Printf("failed to create file %s", err.Error())
			return
		}
		defer out.Close()

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("failed to download file %s", err.Error())
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("bad status: %s", resp.Status)
			return
		}

		bar := progressbar.DefaultBytes(
			resp.ContentLength,
			"downloading db",
		)
		_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
		if err != nil {
			fmt.Printf("failed to copy downloaded file: %s", err.Error())
			return
		}

	}
}

func checkIfLatestRelease() {
	if os.Getenv("INTERLYNK_DISABLE_VERSION_CHECK") != "" {
		return
	}

	client := github.NewClient(nil)
	rr, resp, err := client.Repositories.GetLatestRelease(context.Background(), "interlynk-io", "sbomex")
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		return
	}

	if rr.GetTagName() != version.GetVersionInfo().GitVersion {
		fmt.Printf("\nA new version of sbomex is available %s.\n\n", rr.GetTagName())
	}
}
