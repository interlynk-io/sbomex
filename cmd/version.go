/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"

	version "sigs.k8s.io/release-utils/version"
)

func init() {
	rootCmd.AddCommand(version.Version())
}
