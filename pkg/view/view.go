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

package view

import (
	"context"
	"fmt"
	"os"

	"github.com/interlynk-io/sbomex/pkg/model"
	"github.com/olekukonko/tablewriter"
)

func SearchView(ctx context.Context, sbomex_results []model.SEARCH) {
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
