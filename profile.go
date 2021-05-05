package raiderio

import (
	"context"
	"fmt"
	"github.com/mrburtyyy/go-raiderio/models"
	"net/url"
	"strings"
)

type CharacterProfileFieldName int

const (
	_ CharacterProfileFieldName = iota
	Gear
	Guild
	Covenant
	RaidProgression
	MythicPlusScoresBySeason
	MythicPlusRanks
	MythicPlusRecentRuns
	MythicPlusBestRuns
	MythicPlusHighestLevelRuns
	MythicPlusWeeklyHighestLevelRuns
	MythicPlusPreviousWeeklyHighestLevelRuns
	PreviousMythicPlusRanks
	RaidAchievementMeta
	RaidAchievementCurve
)

func (c CharacterProfileFieldName) String() string {
	rr := []string {
		"",
		"gear",
		"guild",
		"covenant",
		"raid_progression",
		"mythic_plus_scores_by_season",
		"mythic_plus_ranks",
		"mythic_plus_recent_runs",
		"mythic_plus_best_runs",
		"mythic_plus_highest_level_runs",
		"mythic_plus_weekly_highest_level_runs",
		"mythic_plus_previous_weekly_highest_level_runs",
		"previous_mythic_plus_ranks",
		"raid_achievement_meta",
		"raid_achievement_curve",
	}

	return rr[c]
}

type CharacterProfileField struct {
	FieldName CharacterProfileFieldName
	Values    []string
}

func (pf CharacterProfileField) String() string {
	var sb strings.Builder
	sb.WriteString(pf.FieldName.String())
	for _, v := range pf.Values {
		sb.WriteString(fmt.Sprintf(":%s", v))
	}
	return sb.String()
}

func (c *Client) FetchCharacterProfile(ctx context.Context, realm, name string, fields []CharacterProfileField) (*models.Character, error) {
	qp := url.Values{}
	qp.Add("realm", realm)
	qp.Add("name", name)

	f := bindFields(fields)
	if len(f) > 0 {
		qp.Add("fields", f)
	}

	path := fmt.Sprintf("%s/%s", apiBaseUrl, "characters/profile")
	character, err := c.doRequest(ctx, httpGet, path, qp, &models.Character{})
	if err != nil {
		return nil, err
	}

	return character.(*models.Character), nil
}

func bindFields(fields []CharacterProfileField) string {
	var s []string
	for _, f := range fields {
		s = append(s, f.String())
	}
	return strings.Join(s, ",")
}