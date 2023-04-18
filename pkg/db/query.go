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
	"fmt"
	"strings"
)

func addFilter(column, value string, filterExist bool) string {
	var builder strings.Builder
	if filterExist {
		builder.WriteString(" and ")
	} else {
		builder.WriteString(" where ")
	}
	builder.WriteString(column)
	builder.WriteString(" LIKE '%")
	builder.WriteString(value)
	builder.WriteString("%'")
	return builder.String()
}

func search(target string, format, spec, tool string, limit int32) string {

	query := "SELECT id, target, target_version, spec, format, score, creator, creator_version, file_url FROM sboms JOIN scores ON sboms.id = scores.sbom_id "
	filterExist := false
	if format != "" {
		query += addFilter("sboms.format", format, filterExist)
		filterExist = true
	}

	if target != "" {
		query += addFilter("sboms.target || sboms.target_version", target, filterExist)
		filterExist = true
	}

	if spec != "" {
		query += addFilter("sboms.spec", spec, filterExist)
		filterExist = true
	}

	if tool != "" {
		query += addFilter("sboms.creator", tool, filterExist)
		filterExist = true
	}

	query += " limit " + fmt.Sprint(limit)
	return query
}

func url(id int32) string {
	query := fmt.Sprintf("SELECT id, file_url FROM sboms where id = %d", id)
	query += " limit 50"
	return query
}
