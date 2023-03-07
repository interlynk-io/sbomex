/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
)

type SBOMLC struct {
	DB *sql.DB
}

const (
	SqliteDriver     string = "sqlite3"
	SbomlcDataSource string = "./sbomlc.db"
)

func NewSbomlc(driverName, dataSourceName string) (*SBOMLC, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Panic("Failed to connect database", err)
		return nil, err
	}

	return &SBOMLC{
		DB: db,
	}, nil
}

type SBOMS struct {
	ID             int64     `json:"id,omitempty"`
	Source         string    `json:"source,omitempty"`
	Name           string    `json:"name,omitempty"`
	Format         string    `json:"format,omitempty"`
	Spec           string    `json:"spec,omitempty"`
	SpecVersion    string    `json:"spec_version,omitempty"`
	Creator        string    `json:"creator,omitempty"`
	CreatorVersion string    `json:"creator_version,omitempty"`
	Target         string    `json:"target,omitempty"`
	TargetVersion  string    `json:"target_version,omitempty"`
	FileUrl        string    `json:"file_url,omitempty"`
	Visibility     string    `json:"visibility,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type SCORE struct {
	SbomId    int64   `json:"sbom_id,omitempty"`
	Score     float64 `json:"score,omitempty"`
	ScoreJSON string  `json:"score_json,omitempty"`
}

type SEARCH struct {
	SBOMS
	SCORE
}

func (sl *SBOMLC) Search(format, spec, tool string, limit int32) {
	query := "SELECT id, target, target_version, spec, format, score, creator, creator_version, file_url FROM sboms JOIN scores ON sboms.id = scores.sbom_id "

	filterExist := false
	if format != "" {
		query += fmt.Sprintf(" where sboms.format = '%s'", format)
		filterExist = true
	}

	if spec != "" {
		if filterExist {
			query += fmt.Sprintf(" and sboms.spec='%s'", spec)
			filterExist = true
		} else {
			query += fmt.Sprintf(" where sboms.spec='%s'", spec)
			filterExist = true
		}
	}

	if tool != "" {
		if filterExist {
			query += " and sboms.creator LIKE '%" + tool + "%'"
			filterExist = true
		} else {
			query += " where sboms.creator LIKE '%" + tool + "%'"
			filterExist = true
		}
	}

	query += " limit " + fmt.Sprint(limit)

	rows, err := sl.DB.Query(query)
	if err != nil {
		fmt.Printf("query execution failed %v", err)
		return
	}
	defer rows.Close()

	sbomex_results := []SEARCH{}

	for rows.Next() {
		sbomex_result := SEARCH{}
		if err := rows.Scan(&sbomex_result.ID, &sbomex_result.Target, &sbomex_result.TargetVersion, &sbomex_result.Spec, &sbomex_result.Format, &sbomex_result.Score, &sbomex_result.Creator, &sbomex_result.CreatorVersion, &sbomex_result.FileUrl); err != nil {
			fmt.Printf("failed to fetch rows %v", err)
			return
		}
		sbomex_results = append(sbomex_results, sbomex_result)
	}
	searchView(sbomex_results)
}

func searchView(sbomex_results []SEARCH) {
	outDoc := [][]string{}

	for _, s := range sbomex_results {
		l := []string{fmt.Sprint(s.ID), fmt.Sprintf("%s:%s", s.Target, s.TargetVersion), fmt.Sprintf("%.2f", s.Score), fmt.Sprintf("%s-%s", s.Spec, s.Format), fmt.Sprintf("%s-%s", s.Creator, s.CreatorVersion)}
		outDoc = append(outDoc, l)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "TARGET", "QUALITY", "TYPE", "CREATOR"})
	table.SetRowLine(true)
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.AppendBulk(outDoc)
	table.Render()
}
