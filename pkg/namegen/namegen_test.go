package namegen_test

import (
	"strings"
	"testing"

	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/namegen"
)

func TestNameGen(t *testing.T) {
    name := namegen.New()
    name.SetConfig(namegen.Config{
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
    })
    name.Add(namegen.Words, "fear", "world", "pizza", "heroku", "the", "bot", "botprofile", "schmuck")
    name.Add(namegen.Nouns, "cat", "dog", "kitten", "crab", "lizard", "crabcat", "tiger", "lion", "Jimmy")
    name.Add(namegen.Adjectives, "blue", "green", "red", "fear", "pink", "purple", "the", "white", "black", "orange")
    name.Add(namegen.Clans, "-tc", "_Xi_", "[btp]", "cry0", "0x0", "-JjE-")

    lists, err := namegen.LoadMainListCSV("mainlists.csv")
    if err != nil {
        t.Fatal(err)
    }

    name.Consume(lists...)

    run := true
    var testName string
    iterations := 0

    for run {
        testName = name.Generate()
        if strings.Contains(testName, "Fear") &&
            strings.Contains(testName, "Crab") &&
            strings.Contains(testName, "The") &&
            strings.Contains(testName, "Cat") {
            run = false
            t.Log("Iterations: ", iterations, " Name: ", testName)
        }

        if iterations > 1000000 {
            run = false
            t.Log("Iterations: ", iterations, " Name: ", testName, " Error: ", "Timed out")
        }

        iterations++
    }

}
