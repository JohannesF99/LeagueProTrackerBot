package model

type PerksDTO struct {
	StatPerks PerkStatsDTO   `json:"statPerks"`
	Styles    []PerkStyleDTO `json:"styles"`
}

type PerkStatsDTO struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type PerkStyleDTO struct {
	Description string                  `json:"description"`
	Selections  []PerkStyleSelectionDTO `json:"selections"`
	Style       int                     `json:"style"`
}

type PerkStyleSelectionDTO struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}
