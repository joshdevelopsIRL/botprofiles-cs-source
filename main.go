package main

import (
	"os"

	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/bots"
	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/namegen"
	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/rng"
)

const (
	MAIN_APP_CONFIG = "AppConfig.json"
	MAIN_CSV_LIST   = "mainlists.csv"
	OUTPUT_DB       = "botprofile.db"
	BASE_POOL_SIZE  = 12
	MAX_NAME_LENGTH = 60
)

var MainConfig = LoadConfig("./" + MAIN_APP_CONFIG)
var NameCFG = namegen.GetDefaultConfig()
var NameGen = namegen.New()
var RNG = rng.New()

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

	if MainConfig.UseClans {
		if diff > 1 {
			useClans = RNG.RandBool(NameCFG.RandPoolSize / 4)
		} else {
			useClans = RNG.RandBool(NameCFG.RandPoolSize - 2)
		}

	}

	NameCFG.UseClans = useClans
	NameGen.SetConfig(*NameCFG)
}

func main() {
	NameCFG.UseWords = MainConfig.UseWords
	NameCFG.UseNumbers = MainConfig.UseNumbers
	NameCFG.UseNouns = MainConfig.UseNouns
	NameCFG.UseAdjectives = MainConfig.UseAdjectives
	NameCFG.UseClans = MainConfig.UseClans
	NameCFG.UseUppercase = MainConfig.UseUppercase
	NameCFG.UseLowercase = MainConfig.UseLowercase
	NameCFG.MaxNameLength = MainConfig.MaxNameLength
	NameCFG.MaxListLength = MainConfig.MaxListLength
	NameCFG.RandPoolSize = MainConfig.RandPoolSize
	NameCFG.MaxWordsInName = MainConfig.MaxWordsInName

	lists, err := namegen.LoadMainListCSV(MainConfig.CSVNamelistPath)
	if err != nil {
		panic(err)
	}
	NameGen.SetConfig(*NameCFG)
	NameGen.Consume(lists...)

	f, err := os.Create(MainConfig.DBOutputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(bots.GenerateDefaultTemplate())
	f.WriteString(bots.GenerateDifficultyTemplates())
	f.WriteString(bots.GenerateWeaponTemplates())

	generatedAmount := 5

	for diff := range bots.DefaultConfigs {
		for weapon := range bots.WeaponAffinities {

			switch weapon {
			case bots.Shotguns:
				generatedAmount = MainConfig.MaxShotguns
			case bots.SMGs:
				generatedAmount = MainConfig.MaxSMGs
			case bots.Scoped:
				generatedAmount = MainConfig.MaxScoped
			case bots.Autos:
				generatedAmount = MainConfig.MaxAutos
			case bots.Snipers:
				generatedAmount = MainConfig.MaxSnipers
			case bots.Rifles:
				generatedAmount = MainConfig.MaxRifles
			default:
				generatedAmount = 5
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
