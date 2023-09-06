package models

type ElasticViewerConfigYML struct {
	Elasticsearch struct {
		Host              string `yaml:"host"`
		Username          string `yaml:"username"`
		Password          string `yaml:"password"`
		AttemptsToConnect int8   `yaml:"attempts"`
		Interval          int8   `yaml:"attemps"`
		Timeout           int8   `yaml:"timeout"`
	} `yaml:"elasticsearch"`
}
