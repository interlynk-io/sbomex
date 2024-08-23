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

package model

import "time"

const (
	SbomlcDataSource string = ".interlynk-io/sbomex/sqlite3.db"
	DbLocation       string = "https://s3.amazonaws.com/app.interlynk.io/static/db/sbomlc.db"
)

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

type CMDArgs struct {
	Id     int32
	Target string
	Format string
	Spec   string
	Tool   string
	Limit  int32
}
