package bots

type Difficulty int

const (
    Easy Difficulty = iota
    Normal
    Hard
    Expert
)

func (d Difficulty) String() string {
    switch d {
    case Easy:
        return "EASY"
    case Normal:
        return "NORMAL"
    case Hard:
        return "HARD"
    case Expert:
        return "EXPERT"
    }
    return ""
}

func (d Difficulty) Int() int {
    return int(d)
}

func (d Difficulty) TemplateName() string {
    switch d {
    case Easy:
        return "Simple"
    case Normal:
        return "Regular"
    case Hard:
        return "Difficult"
    case Expert:
        return "Broken"
    }
    return ""
}


type WeaponAffinity int

const (
    Rifles WeaponAffinity = iota
    Shotguns
    SMGs
    Snipers
    Scoped
    Autos
)

func (w WeaponAffinity) String() string {
    switch w {
    case Rifles:
        return "Rifles"
    case Shotguns:
        return "Shotguns"
    case SMGs:
        return "SMGs"
    case Snipers:
        return "Snipers"
    case Scoped:
        return "Scoped"
    case Autos:
        return "Autos"
    }
    return ""
}

var WeaponAffinities = map[WeaponAffinity][]string{
    Rifles: {
        "ak47",
        "m4a1",
        "galil",
        "famas",
        "deagle",
        "fiveseven",
        "elite",
    },
    Shotguns: {
        "xm1014",
        "m3",
        "deagle",
        "p228",
    },
    SMGs: {
        "p90",
        "tmp",
        "mac10",
        "mp5navy",
        "ump45",
        "fiveseven",
        "elite",
    },
    Snipers: {
        "awp",
        "m4a1",
        "ak47",
        "scout",
        "deagle",
    },
    Scoped: {
        "aug",
        "sg552",
        "galil",
        "famas",
        "fiveseven",
        "elite",
    },
    Autos: {
        "sg550",
        "g3sg1",
        "m4a1",
        "ak47",
        "scout",
        "deagle",
    },
}
