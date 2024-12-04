package model

type VersionResponse struct {
	Name          string `json:"name"`
	BuildDateTime string `json:"buildDateTime"`
	Version       string `json:"version"`
}
