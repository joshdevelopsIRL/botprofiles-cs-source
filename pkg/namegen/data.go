package namegen

import (
	"encoding/csv"
	"io"
	"os"
)


type ListType int

const (
    Words ListType = iota
    Nouns
    Adjectives
    Clans
)

func (l ListType) String() string {
    switch l {
    case Words:
        return "Words"
    case Nouns:
        return "Nouns"
    case Adjectives:
        return "Adjectives"
    case Clans:
        return "Clans"
    }
    return ""
}

func GetListTypes() []ListType {
    return []ListType{Words, Nouns, Adjectives, Clans}
}

func GetListType(s string) ListType {
    switch s {
    case "Words":
        return Words
    case "Nouns":
        return Nouns
    case "Adjectives":
        return Adjectives
    case "Clans":
        return Clans
    }
    return -1
}

type List struct {
    ListType ListType
    Items []string
}

// This will create multiple lists based upon
// the CSV header fields. It assumes there is
// a header row at the 0th row and the headers are:
// Words, Nouns, Adjectives, Clans
func LoadMainListCSV(fname string) ([]List, error) {
    f, err := os.Open(fname)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    r := csv.NewReader(f)
    r.LazyQuotes = true
    r.TrimLeadingSpace = true

    listPositions := make(map[ListType]int)
    lists := make([]*List, 0)
    for _, ltype := range GetListTypes() {
        listPositions[ltype] = -1
        lists = append(lists, &List{ListType: ltype, Items: make([]string, 0)})
    }



    line := 0
    for {
        row, err := r.Read()
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }

        if line == 0 {
            for i, ltype := range row {
                lt := GetListType(ltype)
                if lt > -1 {
                    listPositions[lt] = i
                }
            }
            line++
            continue
        }

        for listType, index := range listPositions {
            if index < 0 {
                continue
            }
            if len(row) > index {
                for _, list := range lists {
                    if list.ListType == listType {
                        if len(row[index]) > 0 {
                            list.Items = append(list.Items, row[index])
                        }
                    }
                }
            }
        }

        line++
    }

    results := make([]List, 0)
    for _, list := range lists {
        if len(list.Items) > 0 {
            results = append(results, *list)
        }
    }

    return results, nil
}

// This will create a single list based upon
// the given list type. it completely ignores
// the CSV header, in fact, it will add the 0th
// row to the list as well as the subsequent rows
// func NewListFromCSV(fname string, listType ListType) (List, error) {
//     return nil, nil
// }
