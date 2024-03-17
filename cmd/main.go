package main

import (
	"os"

	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/bots"
	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/namegen"
	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/rng"
)

var NameCFG = namegen.GetDefaultConfig()
var NameGen = namegen.New()
var RNG = rng.New()

const BASE_POOL_SIZE = 12

var UsedNames = make([]string, 0)

func isUsed(name string) bool {
    for i := range UsedNames {
        if name == UsedNames[i] {
            return true
        }
    }
    return false
}

func GenerateRandomName() string {
    name := NameGen.Generate()

    if isUsed(name) {
        return GenerateRandomName()
    }

    UsedNames = append(UsedNames, name)
    return name
}

func RandomizeClans(diff int) {
    useClans := false

    if diff > 1 {
        useClans = RNG.RandBool(BASE_POOL_SIZE/4)
    } else {
        useClans = RNG.RandBool(BASE_POOL_SIZE)
    }

    NameCFG.UseClans = useClans

    NameGen.SetConfig(*NameCFG)
}

func main() {
    NameCFG.MaxListLength = 30
    NameCFG.RandPoolSize = BASE_POOL_SIZE
    NameCFG.MaxWordsInName = 7
    NameCFG.MaxNameLength = 24

    lists, err := namegen.LoadMainListCSV("../mainlists.csv")
    if err != nil {
        panic(err)
    }

    NameGen.SetConfig(*NameCFG)
    NameGen.Consume(lists...)

    f, err := os.Create("botprofile.db")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    f.WriteString(bots.GenerateDefaultTemplate())
    f.WriteString(bots.GenerateDifficultyTemplates())
    f.WriteString(bots.GenerateWeaponTemplates())
    
    generatedAmount := 15

    for diff := range bots.DefaultConfigs {
        for weapon := range bots.WeaponAffinities {
            RandomizeClans(diff.Int())

            if weapon == bots.Shotguns ||
               weapon == bots.SMGs ||
               weapon == bots.Scoped {
                generatedAmount = 5
            } else {
                generatedAmount = 15
            }

            for range generatedAmount {
                b := bots.NewProfile(GenerateRandomName(), weapon, diff)
                b.Generate()
                f.WriteString(b.Template())
            }
        }
    }


}
