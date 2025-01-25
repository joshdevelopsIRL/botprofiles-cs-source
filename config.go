package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type AppConfig struct {
	CSVNamelistPath string `json:"csv_namelist_path"`
	DBOutputPath    string `json:"db_output_path"`

	UseWords      bool `json:"use_words"`
	UseNumbers    bool `json:"use_numbers"`
	UseNouns      bool `json:"use_nouns"`
	UseAdjectives bool `json:"use_adjectives"`
	UseClans      bool `json:"use_clans"`
	UseUppercase  bool `json:"use_uppercase"`
	UseLowercase  bool `json:"use_lowercase"`

	MaxNameLength  int `json:"max_name_length"`
	MaxListLength  int `json:"max_list_length"`
	MaxWordsInName int `json:"max_words_in_name"`
	RandPoolSize   int `json:"rand_pool_size"`

	MaxShotguns int `json:"max_shotguns"`
	MaxSMGs     int `json:"max_smgs"`
	MaxScoped   int `json:"max_scoped"`
	MaxAutos    int `json:"max_autos"`
	MaxSnipers  int `json:"max_snipers"`
	MaxRifles   int `json:"max_rifles"`
}

func GetDefaultAppConfig() *AppConfig {
	return &AppConfig{
		CSVNamelistPath: "./" + MAIN_CSV_LIST,
		DBOutputPath:    "./release/" + OUTPUT_DB,

		UseWords:      true,
		UseNumbers:    true,
		UseNouns:      true,
		UseAdjectives: true,
		UseClans:      false,
		UseUppercase:  true,
		UseLowercase:  true,

		MaxNameLength:  40,
		MaxListLength:  10,
		RandPoolSize:   20,
		MaxWordsInName: 3,

		MaxShotguns: 1,
		MaxSMGs:     3,
		MaxScoped:   6,
		MaxAutos:    3,
		MaxSnipers:  6,
		MaxRifles:   9,
	}
}

func validateGun(amt int) bool {
	return amt < 0 || amt >= 100
}

func ValidateAppConfig(cfg *AppConfig) *AppConfig {
	tmp := GetDefaultAppConfig()

	if len(cfg.CSVNamelistPath) == 0 {
		fmt.Println("USING DEFAULT NAME LIST CSV")
		cfg.CSVNamelistPath = tmp.CSVNamelistPath
	}

	if !cfg.UseWords && !cfg.UseNouns && !cfg.UseAdjectives {
		fmt.Println("USING DEFAULT WORDS + NOUNS + ADJ")
		cfg.UseWords = tmp.UseWords
		cfg.UseNouns = tmp.UseNouns
		cfg.UseAdjectives = tmp.UseAdjectives
	}

	if !cfg.UseUppercase && !cfg.UseLowercase {
		fmt.Println("USING DEFAULT UPPER + LOWER CASE")
		cfg.UseUppercase = tmp.UseUppercase
		cfg.UseLowercase = tmp.UseLowercase
	}

	if cfg.MaxNameLength < 3 || cfg.MaxNameLength >= MAX_NAME_LENGTH {
		fmt.Println("USING DEFAULT MAX NAME LEN")
		cfg.MaxNameLength = tmp.MaxNameLength
	}

	if cfg.MaxListLength < 1 || cfg.MaxListLength >= 50 {
		fmt.Println("USING DEFAULT MAX LIST LEN")
		cfg.MaxListLength = tmp.MaxListLength
	}

	if cfg.RandPoolSize < 1 || cfg.RandPoolSize >= 100 {
		fmt.Println("USING DEFAULT RAND POOL SIZE")
		cfg.RandPoolSize = tmp.RandPoolSize
	}

	if cfg.MaxWordsInName < 1 || cfg.MaxWordsInName >= 10 {
		fmt.Println("USING DEFAULT MAX WORDS IN NAME")
		cfg.MaxWordsInName = tmp.MaxWordsInName
	}

	if validateGun(cfg.MaxShotguns) {
		fmt.Println("USING DEFAULT MAX SHOTGUNS")
		cfg.MaxShotguns = tmp.MaxShotguns
	}

	if validateGun(cfg.MaxSMGs) {
		fmt.Println("USING DEFAULT MAX SMGs")
		cfg.MaxSMGs = tmp.MaxSMGs
	}

	if validateGun(cfg.MaxScoped) {
		fmt.Println("USING DEFAULT MAX SCOPED")
		cfg.MaxScoped = tmp.MaxScoped
	}

	if validateGun(cfg.MaxAutos) {
		fmt.Println("USING DEFAULT MAX AUTOS")
		cfg.MaxAutos = tmp.MaxAutos
	}

	if validateGun(cfg.MaxSnipers) {
		fmt.Println("USING DEFAULT MAX SNIPERS")
		cfg.MaxSnipers = tmp.MaxSnipers
	}

	if validateGun(cfg.MaxRifles) {
		fmt.Println("USING DEFAULT MAX RIFLES")
		cfg.MaxRifles = tmp.MaxRifles
	}

	return cfg
}

func LoadConfig(jsonPath string) *AppConfig {
	if len(jsonPath) > 0 {
		if !strings.Contains(jsonPath, MAIN_APP_CONFIG) {
			if !strings.HasSuffix(jsonPath, "/") {
				jsonPath += "/"
			}
			jsonPath += MAIN_APP_CONFIG
		}
	} else {
		jsonPath = "../AppConfig.json"
	}

	conf := GetDefaultAppConfig()

	b, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("USING DEFAULT CONFIG, UNABLE TO READ JSON PATH ::", err)
		return conf
	}

	var tmp AppConfig
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		fmt.Println("USING DEFAULT CONFIG, UNABLE TO PARSE JSON ::", err)
		return conf
	}

	conf.CSVNamelistPath = tmp.CSVNamelistPath
	conf.DBOutputPath = tmp.DBOutputPath
	conf.UseWords = tmp.UseWords
	conf.UseNumbers = tmp.UseNumbers
	conf.UseNouns = tmp.UseNouns
	conf.UseAdjectives = tmp.UseAdjectives
	conf.UseClans = tmp.UseClans
	conf.UseUppercase = tmp.UseUppercase
	conf.UseLowercase = tmp.UseLowercase
	conf.MaxNameLength = tmp.MaxNameLength
	conf.MaxListLength = tmp.MaxListLength
	conf.RandPoolSize = tmp.RandPoolSize
	conf.MaxWordsInName = tmp.MaxWordsInName
	conf.MaxShotguns = tmp.MaxShotguns
	conf.MaxSMGs = tmp.MaxSMGs
	conf.MaxScoped = tmp.MaxScoped
	conf.MaxAutos = tmp.MaxAutos
	conf.MaxSnipers = tmp.MaxSnipers
	conf.MaxRifles = tmp.MaxRifles

	return ValidateAppConfig(conf)
}
