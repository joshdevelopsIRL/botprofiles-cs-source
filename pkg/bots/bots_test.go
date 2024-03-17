package bots_test

import (
	"testing"

	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/bots"
	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/namegen"
)

func TestBots(t *testing.T) {
    lists, err := namegen.LoadMainListCSV("../namegen/mainlists.csv")
    if err != nil {
        t.Fatal(err)
    }
    gen := namegen.New()
    gen.Consume(lists...)

    nameCfg := &namegen.Config{
        MaxNameLength: 24,
        UseWords: true,
        UseNumbers: true,
        UseNouns: true,
        UseAdjectives: true,
        UseClans: true,
        UseUppercase: true,
        UseLowercase: true,
        MaxListLength: 5,
        RandPoolSize: 30,
        MaxWordsInName: 5,
    }


    for diff := range bots.DefaultConfigs {
        if diff.Int() > 1 {
            nameCfg.UseClans = true
        } else {
            nameCfg.UseClans = false
        }
        for weapon := range bots.WeaponAffinities {
            gen.SetConfig(*nameCfg)

            b := bots.NewProfile(gen.Generate(), weapon, diff)
            b.Generate()
            t.Log(b.String())
            t.Log(b.Template())
        }
    }
}
