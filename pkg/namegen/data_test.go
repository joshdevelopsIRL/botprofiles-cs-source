package namegen_test

import (
	"testing"

	"github.com/joshdevelopsIRL/botprofiles-cs-source/pkg/namegen"
)

func TestData(t *testing.T) {
    fname := "mainlists.csv"
    lists, err := namegen.LoadMainListCSV(fname)
    if err != nil {
        t.Fatal(err)
    }
    for _, l := range lists {
        t.Log(l)
    }
}
