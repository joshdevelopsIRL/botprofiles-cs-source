package bots

type Config struct {
	MaxSkill        int
	MinSkill        int
	MaxAggression   int
	MinAggression   int
	MaxReactionTime float64
	MinReactionTime float64
	MaxAttackDelay  float64
	MinAttackDelay  float64
	MaxTeamwork     int
	MinTeamwork     int
	MaxVoicePitch   int
	MinVoicePitch   int
}

func GetDefaultConfig() *Config {
	return &Config{
		MaxSkill:        50,
		MinSkill:        10,
		MaxAggression:   50,
		MinAggression:   10,
		MaxReactionTime: 1.0,
		MinReactionTime: 0.3,
		MaxAttackDelay:  1.0,
		MinAttackDelay:  0.3,
		MaxTeamwork:     50,
		MinTeamwork:     10,
		MaxVoicePitch:   200,
		MinVoicePitch:   100,
	}
}

var DefaultConfigs = map[Difficulty]Config{
	Easy: {
		MaxSkill:        50,
		MinSkill:        10,
		MaxAggression:   50,
		MinAggression:   10,
		MaxReactionTime: 3.0,
		MinReactionTime: 1.0,
		MaxAttackDelay:  3.0,
		MinAttackDelay:  1.0,
		MaxTeamwork:     50,
		MinTeamwork:     10,
		MaxVoicePitch:   200,
		MinVoicePitch:   100,
	},
	Normal: {
		MaxSkill:        70,
		MinSkill:        30,
		MaxAggression:   80,
		MinAggression:   40,
		MaxReactionTime: 2.0,
		MinReactionTime: 0.5,
		MaxAttackDelay:  1.0,
		MinAttackDelay:  0.1,
		MaxTeamwork:     80,
		MinTeamwork:     30,
		MaxVoicePitch:   200,
		MinVoicePitch:   100,
	},
	Hard: {
		MaxSkill:        85,
		MinSkill:        65,
		MaxAggression:   90,
		MinAggression:   60,
		MaxReactionTime: 1.3,
		MinReactionTime: 0.5,
		MaxAttackDelay:  1.3,
		MinAttackDelay:  0.5,
		MaxTeamwork:     90,
		MinTeamwork:     70,
		MaxVoicePitch:   200,
		MinVoicePitch:   100,
	},
	Expert: {
		MaxSkill:        100,
		MinSkill:        85,
		MaxAggression:   100,
		MinAggression:   85,
		MaxReactionTime: 1.0,
		MinReactionTime: 0.1,
		MaxAttackDelay:  1.0,
		MinAttackDelay:  0.1,
		MaxTeamwork:     100,
		MinTeamwork:     85,
		MaxVoicePitch:   100,
		MinVoicePitch:   200,
	},
}
