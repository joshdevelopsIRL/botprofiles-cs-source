package bots

const VERSION = "0.7"

type Config struct {
    MaxSkill int
    MinSkill int
    MaxAggression int
    MinAggression int
    MaxReactionTime float64
    MinReactionTime float64
    MaxAttackDelay float64
    MinAttackDelay float64
    MaxTeamwork int
    MinTeamwork int
    MaxVoicePitch int
    MinVoicePitch int
}

func GetDefaultConfig() *Config {
    return &Config{
        MaxSkill: 50,
        MinSkill: 10,
        MaxAggression: 50,
        MinAggression: 10,
        MaxReactionTime: 1.0,
        MinReactionTime: 0.3,
        MaxAttackDelay: 1.0,
        MinAttackDelay: 0.3,
        MaxTeamwork: 50,
        MinTeamwork: 10,
        MaxVoicePitch: 200,
        MinVoicePitch: 100,
    }
}

var DefaultConfigs = map[Difficulty]Config{
    Easy: {
        MaxSkill: 50,
        MinSkill: 10,
        MaxAggression: 50,
        MinAggression: 10,
        MaxReactionTime: 1.0,
        MinReactionTime: 0.3,
        MaxAttackDelay: 1.0,
        MinAttackDelay: 0.3,
        MaxTeamwork: 50,
        MinTeamwork: 10,
        MaxVoicePitch: 200,
        MinVoicePitch: 100,
    },
    Normal: {
        MaxSkill: 70,
        MinSkill: 40,
        MaxAggression: 70,
        MinAggression: 40,
        MaxReactionTime: 0.7,
        MinReactionTime: 0.1,
        MaxAttackDelay: 0.7,
        MinAttackDelay: 0.1,
        MaxTeamwork: 70,
        MinTeamwork: 40,
        MaxVoicePitch: 200,
        MinVoicePitch: 100,
    },
    Hard: {
        MaxSkill: 90,
        MinSkill: 60,
        MaxAggression: 90,
        MinAggression: 60,
        MaxReactionTime: 0.3,
        MinReactionTime: 0.01,
        MaxAttackDelay: 0.03,
        MinAttackDelay: 0.01,
        MaxTeamwork: 100,
        MinTeamwork: 70,
        MaxVoicePitch: 200,
        MinVoicePitch: 100,
    },
    Expert: {
        MaxSkill: 100,
        MinSkill: 90,
        MaxAggression: 100,
        MinAggression: 90,
        MaxReactionTime: 0.1,
        MinReactionTime: 0.01,
        MaxAttackDelay: 0.1,
        MinAttackDelay: 0.01,
        MaxTeamwork: 100,
        MinTeamwork: 90,
        MaxVoicePitch: 100,
        MinVoicePitch: 200,
    },
}
