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
		MaxSkill:        70,
		MinSkill:        10,
		MaxAggression:   70,
		MinAggression:   10,
		MaxReactionTime: 1.0,
		MinReactionTime: 0.3,
		MaxAttackDelay:  1.0,
		MinAttackDelay:  0.3,
		MaxTeamwork:     70,
		MinTeamwork:     30,
		MaxVoicePitch:   200,
		MinVoicePitch:   100,
	},
	Normal: {
		MaxSkill:        90,
		MinSkill:        10,
		MaxAggression:   90,
		MinAggression:   10,
		MaxReactionTime: 1.0,
		MinReactionTime: 0.1,
		MaxAttackDelay:  1.0,
		MinAttackDelay:  0.1,
		MaxTeamwork:     90,
		MinTeamwork:     10,
		MaxVoicePitch:   200,
		MinVoicePitch:   100,
	},
	Hard: {
		MaxSkill:        90,
		MinSkill:        60,
		MaxAggression:   90,
		MinAggression:   60,
		MaxReactionTime: 0.7,
		MinReactionTime: 0.1,
		MaxAttackDelay:  0.7,
		MinAttackDelay:  0.1,
		MaxTeamwork:     90,
		MinTeamwork:     60,
		MaxVoicePitch:   200,
		MinVoicePitch:   100,
	},
	Expert: {
		MaxSkill:        100,
		MinSkill:        80,
		MaxAggression:   100,
		MinAggression:   80,
		MaxReactionTime: 0.4,
		MinReactionTime: 0.1,
		MaxAttackDelay:  0.4,
		MinAttackDelay:  0.1,
		MaxTeamwork:     100,
		MinTeamwork:     80,
		MaxVoicePitch:   200,
		MinVoicePitch:   100,
	},
}
