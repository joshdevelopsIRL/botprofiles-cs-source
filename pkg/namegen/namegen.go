package namegen

import (
	"strings"

	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/rng"
)

type Config struct {
    MaxNameLength int
    WantedSeparator string
    UseWords bool
    UseNumbers bool
    UseNouns bool
    UseAdjectives bool
    UseClans bool
    UseUppercase bool
    UseLowercase bool
    MaxListLength int
    RandPoolSize int
    MaxWordsInName int
}

func GetDefaultConfig() *Config {
    return &Config{
        MaxNameLength: 18,
        UseWords: true,
        UseNumbers: true,
        UseNouns: true,
        UseAdjectives: true,
        UseClans: false,
        UseUppercase: true,
        UseLowercase: true,
        MaxListLength: 10,
        RandPoolSize: 20,
        MaxWordsInName: 3,
    }
}

var DefaultConfig = Config{
    MaxNameLength: 18,
    UseWords: true,
    UseNumbers: true,
    UseNouns: true,
    UseAdjectives: true,
    UseClans: false,
    UseUppercase: true,
    UseLowercase: true,
    MaxListLength: 10,
    RandPoolSize: 20,
    MaxWordsInName: 3,
}

var NUMBERS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var SEPARATORS = []string{" ", "", " ", "-", " ", "_", " ", "."}

type Generator struct {
    r rng.RNG
    wordList []string
    nounList []string
    adjectiveList []string
    clanList []string
    config Config
}

func New() *Generator {
    return &Generator{
        r: rng.New(),
        config: DefaultConfig,
    }
}

func (g *Generator) Consume(lists ...List) {
    for _, list := range lists {
        g.Add(list.ListType, list.Items...)
    }
}

func (g *Generator) Add(list ListType, items ...string) {
    switch list {
    case Words:
        g.wordList = append(g.wordList, items...)
    case Nouns:
        g.nounList = append(g.nounList, items...)
    case Adjectives:
        g.adjectiveList = append(g.adjectiveList, items...)
    case Clans:
        g.clanList = append(g.clanList, items...)
    }
}

func (g *Generator) SetConfig(config Config) {
    g.config = config
}

func (g *Generator) getRandList(list ListType) []string {
    result := []string{}

    for range g.config.MaxListLength {
        var w string

        switch list {
        case Words:
            g.r.Shuffle(g.wordList)
            w = g.r.RandChoice(g.wordList)
        case Nouns:
            g.r.Shuffle(g.nounList)
            w = g.r.RandChoice(g.nounList)
        case Adjectives:
            g.r.Shuffle(g.adjectiveList)
            w = g.r.RandChoice(g.adjectiveList)
        default:
            w = "NIL" + g.r.RandChoice(NUMBERS)
        }

        if g.config.UseNumbers && g.r.RandBool(g.config.RandPoolSize) {
            w += g.r.RandChoice(NUMBERS)
        }

        w = strings.ToUpper(w[0:1]) + w[1:]

        if g.config.UseUppercase && g.r.RandBool(g.config.RandPoolSize) {
            w = strings.ToUpper(w)
        }
        if g.config.UseLowercase && g.r.RandBool(g.config.RandPoolSize) {
            w = strings.ToLower(w)
        }

        result = append(result, w)
    }

    return result
}

func (g *Generator) Generate() string {
    wordPool := []string{}

    clan := ""
    if len(g.clanList) > 0 {
        g.r.Shuffle(g.clanList)
        clan = g.r.RandChoice(g.clanList)
    }

    sep := g.r.RandChoice(SEPARATORS)
    if g.config.WantedSeparator != "" {
        sep = g.config.WantedSeparator
    }

    if g.config.UseWords && len(g.wordList) > 0 {
        wordPool = append(wordPool, g.getRandList(Words)...)
    }
    if g.config.UseNouns && len(g.nounList) > 0 {
        wordPool = append(wordPool, g.getRandList(Nouns)...)
    }
    if g.config.UseAdjectives && len(g.adjectiveList) > 0 {
        wordPool = append(wordPool, g.getRandList(Adjectives)...)
    }

    maxWords := g.r.RandInt(1, g.config.MaxWordsInName)
    baseName := []string{}

    for range maxWords {
        baseName = append(baseName, g.r.RandChoice(wordPool))
    }

    if g.config.UseClans && len(clan) > 0 {
        clan = clan + sep
    } else {
        clan = ""
    }

    name := strings.Join(baseName, sep)
    if len(name) > g.config.MaxNameLength {
        name = name[0:g.config.MaxNameLength]
    }

    return clan + name
}
