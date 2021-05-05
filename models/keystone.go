package models

// KeystoneRun represents information about a singular Mythic keystone
// run that has been completed.
type KeystoneRun struct {
	Dungeon             string  `json:"dungeon"`
	ShortName           string  `json:"short_name"`
	MythicLevel         int     `json:"mythic_level"`
	CompletedAt         string  `json:"completed_at"`
	ClearTimeMs         int     `json:"clear_time_ms"`
	NumKeystoneUpgrades int     `json:"num_keystone_upgrades"`
	Score               float64 `json:"score"`
	URL                 string  `json:"url"`
}