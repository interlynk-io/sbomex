/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package view

import (
	"context"
	"fmt"
	"os"

	"github.com/interlynk-io/sbomex/pkg/model"
	"github.com/olekukonko/tablewriter"
)
	

func SearchView(ctx context.Context,  sbomex_results []model.SEARCH) {
	outDoc := [][]string{}

	for _, s := range sbomex_results {
		l := []string{fmt.Sprint(s.ID), fmt.Sprintf("%s:%s", s.Target, s.TargetVersion), fmt.Sprintf("%.2f", s.Score), fmt.Sprintf("%s-%s", s.Spec, s.Format), fmt.Sprintf("%s-%s", s.Creator, s.CreatorVersion)}
		outDoc = append(outDoc, l)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ID", "TARGET", "QUALITY", "TYPE", "CREATOR"})
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.AppendBulk(outDoc)
	table.Render()
}
