package models

type Pass struct {
	NewPass     string `json:"newPass"`
	CurrentPass string `json:"currentPass"`
}
