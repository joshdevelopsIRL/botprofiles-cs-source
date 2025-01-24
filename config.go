package main

const MAX_NAME_LENGTH = 60

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
}

func LoadConfig(jsonPath string) *AppConfig {
	// TODO: Do the json loading thingy and output the conf struct and
	// merge it to overwrite the DefaultConfig. Must not break the
	// MAX_NAME_LENGTH on the MaxNameLength attribute. if GT then force MAX -1;
	return &AppConfig{}
}
