package model

type VersionModel struct {
	Name          string `json:"name"`
	BuildDateTime string `json:"buildDateTime"`
	Version       string `json:"version"`
}
