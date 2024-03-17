package rng

import (
	"math/rand"
	"time"
)

type RNG struct {
    Seed int64
    r *rand.Rand
}

func New() RNG {
    seed := time.Now().UnixNano()
    return RNG{
        Seed: seed,
        r: rand.New(rand.NewSource(seed)),
    }
}

func (r RNG) RandInt(min, max int) int {
    if max < 2 {
        max = 2
    }
    if min > max {
        min, max = max, min
    }
    if min == max {
        return min
    }
    return min + r.r.Intn(max-min+1)
}

func (r RNG) RandFloat(min, max float64) float64 {
    if max < min {
        min, max = max, min
    }
    if min == max {
        return min
    }
    return min + r.r.Float64()*(max-min) 
}

func (r RNG) RandBool(pool int) bool {
    if pool < 2 {
        pool = 2
    }
    return r.r.Intn(pool) == 1
}

func (r RNG) RandChoice(choices []string) string {
    if len(choices) == 0 {
        return ""
    }
    if len(choices) == 1 {
        return choices[0]
    }
    return choices[r.r.Intn(len(choices))]
}

func (r RNG) Shuffle(list []string) {
    if len(list) < 3 {
        return
    }
    r.r.Shuffle(len(list), func(i, j int) {
        list[i], list[j] = list[j], list[i]
    })
}
