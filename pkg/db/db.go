/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

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
	sqliteDriver     string = "sqlite3"
	sbomlcDataSource string = "./sbomlc.db"
)

func NewSbomlc() (*SBOMLC, error) {
	db, err := sql.Open(sqliteDriver, sbomlcDataSource)
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
