package main

import (
	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/bots"
	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/namegen"
	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/rng"
)

func main() {
    lists, err := namegen.LoadMainListCSV("mainlists.csv")
    if err != nil {
        panic(err)
    }

    r := rng.New()

    nameCFG := namegen.GetDefaultConfig()
    nameCFG.MaxListLength = 30
    nameCFG.RandPoolSize = 10
    nameCFG.MaxWordsInName = 7
    nameCFG.MaxNameLength = 24

    nameMaker := namegen.New()
    nameMaker.SetConfig(*nameCFG)
    nameMaker.Consume(lists...)

    pool := nameCFG.RandPoolSize

    for diff := range bots.DefaultConfigs {
        if diff.Int() > 1 {
            pool = pool / 2
        } else {
            pool = nameCFG.RandPoolSize
        }


        for weapon := range bots.WeaponAffinities {
        }
    }


}
