package dto

type TraitRarityRequest struct {
	CollectionId string       `json:"collectionId" binding:"required"`
	Properties   []PropertyKV `json:"properties" binding:"required,dive"`
}

type PropertyKV struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type TraitRarityResponse struct {
	Traits []ExtendedTraitProperty `json:"traits"`
}

type ExtendedTraitProperty struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Rarity string `json:"rarity"`
}
