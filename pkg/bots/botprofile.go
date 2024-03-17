package bots

import (
	"fmt"
	"strings"

	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/rng"
)

type BotProfile struct {
    Name string
    Skill int
    Aggression int
    ReactionTime float64
    AttackDelay float64
    Teamwork int
    WeaponAffinity WeaponAffinity
    Difficulty Difficulty
    VoicePitch int
    Skin int
    r rng.RNG
    config Config
}

func NewProfile(name string, weaponAffinity WeaponAffinity, difficulty Difficulty) *BotProfile {
    b := &BotProfile{
        Name: name,
        WeaponAffinity: weaponAffinity,
        Difficulty: difficulty,
        r: rng.New(),
        config: DefaultConfigs[difficulty],
    }

    return b
}

func (b *BotProfile) SetConfig(config Config) {
    b.config = config
}

func (b *BotProfile) Generate() {
    b.Skill = b.r.RandInt(b.config.MinSkill, b.config.MaxSkill)
    b.Aggression = b.r.RandInt(b.config.MinAggression, b.config.MaxAggression)
    b.ReactionTime = b.r.RandFloat(b.config.MinReactionTime, b.config.MaxReactionTime)
    b.AttackDelay = b.r.RandFloat(b.config.MinAttackDelay, b.config.MaxAttackDelay)
    b.Teamwork = b.r.RandInt(b.config.MinTeamwork, b.config.MaxTeamwork)
    b.VoicePitch = b.r.RandInt(b.config.MinVoicePitch, b.config.MaxVoicePitch)
    b.Skin = b.r.RandInt(0, 3)
}

func (b *BotProfile) String() string {
    return fmt.Sprintf("Name: \"%s\" - %s - %s\n", b.Name, b.WeaponAffinity.String(), b.Difficulty.String())
}

func (b *BotProfile) Template() string {
    r := &strings.Builder{}

    r.WriteString(b.Difficulty.TemplateName() + "+" + b.WeaponAffinity.String())
    r.WriteString(" ")
    r.WriteString("\"" + b.Name + "\"")
    r.WriteString("\n")

    r.WriteString(fmt.Sprintf("\tSkill: %d\n", b.Skill))
    r.WriteString(fmt.Sprintf("\tAggression: %d\n", b.Aggression))
    r.WriteString(fmt.Sprintf("\tReactionTime: %.4f\n", b.ReactionTime))
    r.WriteString(fmt.Sprintf("\tAttackDelay: %.4f\n", b.AttackDelay))
    r.WriteString(fmt.Sprintf("\tTeamwork: %d\n", b.Teamwork))
    r.WriteString(fmt.Sprintf("\tVoicePitch: %d\n", b.VoicePitch))
    r.WriteString(fmt.Sprintf("\tSkin: %d\n", b.Skin))

    r.WriteString("End\n\n")
    return r.String()
}
