package models

type RaidName string

const (
	TheEmeraldNightmare     = "the-emerald-nighmare"
	TrialOfValor            = "trial-of-valor"
	TheNighthold            = "the-nighthold"
	TombOfSargeras          = "tomb-of-sargeras"
	AntorusTheBurningThrone = "antorus-the-burning-throne"
	Uldir                   = "uldir"
	BattleOfDazaralor       = "battle-of-dazaralor"
	CrucibleOfStorms        = "crucible-of-storms"
	TheEternalPalace        = "the-eternal-palace"
	NyalothaTheWakingCity   = "nyalotha-the-waking-city"
	CastleNathria           = "castle-nathria"
	SanctumOfDomination     = "sanctum-of-domination"
)

type RaidProgression struct {
	Summary string `json:"summary"`
	TotalBosses int `json:"total_bosses"`
	NormalBossesKilled int `json:"normal_bosses_killed"`
	HeroicBossesKilled int `json:"heroic_bosses_killed"`
	MythicBossesKilled int `json:"Mythic_bosses_killed"`
}