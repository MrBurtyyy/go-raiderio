package models

type Gear struct {
	ItemLevelEquipped int `json:"item_level_equipped"`
	ItemLevelTotal int `json:"item_level_total"`
	ArtifactTraits float32 `json:"artifact_traits"`
}