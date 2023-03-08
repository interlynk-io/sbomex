/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package db

import (
	"fmt"
)

func search(format, spec, tool string, limit int32) string {
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
	return query
}
