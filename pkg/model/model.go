/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package model

import "time"

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