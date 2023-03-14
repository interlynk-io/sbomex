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

package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/interlynk-io/sbomex/pkg/model"
	"github.com/interlynk-io/sbomex/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

type SBOMLC struct {
	DB *sql.DB
}

const (
	sqliteDriver string = "sqlite3"
)

func NewSbomlc() (*SBOMLC, error) {
	db, err := sql.Open(sqliteDriver, model.SbomlcDataSource)
	if err != nil {
		log.Panic("Failed to connect database", err)
		return nil, err
	}

	return &SBOMLC{
		DB: db,
	}, nil
}

func (sl *SBOMLC) Search(ca *model.CMDArgs) []model.SEARCH {
	sbomex_results := []model.SEARCH{}

	rows, err := sl.DB.Query(search(ca.Format, ca.Spec, ca.Tool, ca.Limit))
	if err != nil {
		fmt.Printf("query execution failed %v", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		sbomex_result := model.SEARCH{}
		if err := rows.Scan(&sbomex_result.ID, &sbomex_result.Target, &sbomex_result.TargetVersion, &sbomex_result.Spec, &sbomex_result.Format, &sbomex_result.Score, &sbomex_result.Creator, &sbomex_result.CreatorVersion, &sbomex_result.FileUrl); err != nil {
			fmt.Printf("failed to fetch rows %v", err)
			return nil
		}
		sbomex_results = append(sbomex_results, sbomex_result)
	}
	return sbomex_results
}

func (sl *SBOMLC) Url(ca *model.CMDArgs) string {
	sbomex_results := []*model.SBOMS{}
	var furl string
	rows, err := sl.DB.Query(url(ca.Id))
	if err != nil {
		fmt.Printf("query execution failed %v", err)
		return furl
	}
	defer rows.Close()

	for rows.Next() {
		sbomex_result := &model.SBOMS{}
		if err := rows.Scan(&sbomex_result.ID, &sbomex_result.FileUrl); err != nil {
			fmt.Printf("failed to fetch rows %v", err)
			return furl
		}
		sbomex_results = append(sbomex_results, sbomex_result)
	}
	if len(sbomex_results) == 0 {
		fmt.Println("no record found")
		return furl
	}
	return sbomex_results[utils.RandomPick(0, len(sbomex_results))].FileUrl
}
