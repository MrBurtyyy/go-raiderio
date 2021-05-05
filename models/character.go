package models

// Character represents the information about a specific World of Warcraft character.
type Character struct {
	Name                      string                     `json:"name"`
	Race                      string                     `json:"race"`
	Class                     string                     `json:"class"`
	CurrentSpec               string                     `json:"active_spec_name"`
	CurrentRole               string                     `json:"active_spec_role"`
	Gender                    string                     `json:"gender"`
	Faction                   string                     `json:"faction"`
	Region                    string                     `json:"Region"`
	Realm                     string                     `json:"realm"`
	ProfileUrl                string                     `json:"profile_url"`
	Gear                      Gear                       `json:"gear"`
	WeeklyHighestRuns         []KeystoneRun              `json:"mythic_plus_weekly_highest_level_runs"`
	PreviousWeeklyHighestRuns []KeystoneRun              `json:"mythic_plus_previous_weekly_highest_level_runs"`
	RaidProgression           map[RaidName]RaidProgression `json:"raid_progression"`
}

// RaidProgressionForRaid returns the RaidProgression for the named raid e.g. 'sanctum-of-domination'
func (c Character) RaidProgressionForRaid(name RaidName) RaidProgression {
	return c.RaidProgression[name]
}

// NumberOfKeysCompletedInPreviousWeek returns the number of keys in the previous week, that the
// character completed above a certain level.  This is limited to the 10 highest keys performed in
// the previous week.
func (c Character) NumberOfKeysCompletedInCurrentWeek(minLevel int) int {
	return numberOfKeys(c.WeeklyHighestRuns, minLevel)
}

// NumberOfKeysCompletedInPreviousWeek returns the number of keys in the previous week, that the
// character completed above a certain level.  This is limited to the 10 highest keys performed in
// the previous week.
func (c Character) NumberOfKeysCompletedInPreviousWeek(minLevel int) int {
	return numberOfKeys(c.PreviousWeeklyHighestRuns, minLevel)
}

func numberOfKeys(r []KeystoneRun, m int) int {
	filtered := filterKeys(r, func(run KeystoneRun) bool {
		return run.MythicLevel >= m
	})
	return len(filtered)
}

func filterKeys(vs []KeystoneRun, f func(KeystoneRun) bool) []KeystoneRun {
	vsf := make([]KeystoneRun, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
