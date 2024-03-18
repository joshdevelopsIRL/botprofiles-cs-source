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
        useClans = RNG.RandBool(BASE_POOL_SIZE-2)
    }

    NameCFG.UseClans = useClans

    NameGen.SetConfig(*NameCFG)
}

func main() {
    NameCFG.MaxListLength = 50
    NameCFG.RandPoolSize = BASE_POOL_SIZE
    NameCFG.MaxWordsInName = 5
    NameCFG.MaxNameLength = 18

    lists, err := namegen.LoadMainListCSV("../mainlists.csv")
    if err != nil {
        panic(err)
    }

    NameGen.SetConfig(*NameCFG)
    NameGen.Consume(lists...)

    f, err := os.Create("../botprofile.db")
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

            if weapon == bots.Shotguns ||
               weapon == bots.SMGs { 
                generatedAmount = 5
            } else if weapon == bots.Scoped ||
                      weapon == bots.Autos {
                generatedAmount = 10
            } else if weapon == bots.Snipers {
                generatedAmount = 15
            } else {
                generatedAmount = 25
            }

            for range generatedAmount {
                RandomizeClans(diff.Int())
                b := bots.NewProfile(GenerateRandomName(), weapon, diff)
                b.Generate()
                f.WriteString(b.Template())
            }
        }
    }


}
